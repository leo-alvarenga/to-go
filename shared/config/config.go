package config

import (
	"fmt"
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

func (cfg *ConfigValue) LoadFromYaml(filename string) {
	file, err := os.ReadFile(filename)

	if err != nil {
		file, err = yaml.Marshal(cfg)

		f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("open")
		}

		_, err = f.Write(file)
		if err != nil {
			fmt.Println("w")
		}

		f.Close()

		return
	}

	err = yaml.Unmarshal(file, cfg)

	if err != nil {
		cfg.New()
		return
	}
}
