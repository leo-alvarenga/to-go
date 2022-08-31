package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type ConfigValue struct {
	FirstExec  bool `yaml:"firstExec"`
	MaxLineLen int  `yaml:"maxLineLen"`
}

func (cfg *ConfigValue) New() {
	cfg = new(ConfigValue)

	cfg.SetDefault()
}

func (cfg *ConfigValue) SetDefault() {
	cfg.FirstExec = true
	cfg.MaxLineLen = DefaultMaxLineLen
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

	fmt.Println("ok")
	if err != nil {
		cfg.SetDefault()
		return
	}
}
