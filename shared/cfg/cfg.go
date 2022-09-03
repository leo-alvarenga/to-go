package cfg

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type ColorScheme struct {
	Title   string
	Accent  string
	Error   string
	Warning string
	Success string

	Reset string
}

type ConfigValue struct {
	FirstExec  bool `yaml:"firstExec"`
	MaxLineLen int  `yaml:"maxLineLen"`
	Colors     ColorScheme
}

func (cfg *ConfigValue) New() {
	cfg.FirstExec = true
	cfg.MaxLineLen = DefaultMaxLineLen

	cfg.Colors.Accent += DefaultColors.Accent
	cfg.Colors.Title += DefaultColors.Title
	cfg.Colors.Error += DefaultColors.Error
	cfg.Colors.Warning += DefaultColors.Warning
	cfg.Colors.Success += DefaultColors.Success
	cfg.Colors.Reset += DefaultColors.Reset
}

func (cfg *ConfigValue) LoadFromYaml(filename string) error {
	file, err := os.ReadFile(filename)

	if err != nil {
		file, err = yaml.Marshal(cfg)

		if err != nil {
			return errors.New("Could not save config.")
		}

		f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return errors.New("Could not save config.")
		}

		_, err = f.Write(file)
		if err != nil {
			return errors.New("Could not save config.")
		}

		f.Close()
	}

	err = yaml.Unmarshal(file, cfg)

	if err != nil {
		cfg.New()
	}

	return nil
}
