package cfg

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

const colorReset string = "\033[0m"

type PriorityColor struct {
	High   string `yaml:"high"`
	Medium string `yaml:"medium"`
	Low    string `yaml:"low"`
}

type StatusColor struct {
	Pending string `yaml:"pending"`
	Doing   string `yaml:"doing"`
	Done    string `yaml:"done"`
}

type ColorScheme struct {
	Priority  PriorityColor `yaml:"priority"`
	Status    StatusColor   `yaml:"status"`
	Attention string        `yaml:"attention"`
	Success   string        `yaml:"success"`
	Warning   string        `yaml:"warning"`
	Error     string        `yaml:"error"`
	Reset     string        `yaml:"reset"`
}

type ConfigValue struct {
	UseUnicode     bool        `yaml:"useUnicode"`
	Storage        string      `yaml:"storage"`
	TasksOlderThan int         `yaml:"tasksOlderThan"`
	Colors         ColorScheme `yaml:"colors"`
}

func (cfg *ConfigValue) New() {
	cfg.UseUnicode = true
	cfg.Storage = "sqlite"
	cfg.TasksOlderThan = 10
	cfg.Colors = ColorScheme{
		Priority: PriorityColor{
			High:   "red",
			Medium: "yellow",
			Low:    "green",
		},

		Status: StatusColor{
			Pending: "yellow",
			Doing:   "blue",
			Done:    "green",
		},
		Attention: "purple",
		Success:   "green",
		Warning:   "yellow",
		Error:     "red",
		Reset:     colorReset,
	}
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

func (c *ConfigValue) UseSQLite() bool {
	return c.Storage == "sqlite"
}
