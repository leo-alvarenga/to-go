package ng

import (
	"github.com/leo-alvarenga/to-go/shared/cfg"
	"github.com/leo-alvarenga/to-go/shared/task"
)

// The name of the YAML file used for storing tasks with unassigned priority
var taskFilenames = [3]string{
	"high_p.yaml",
	"medium_p.yaml",
	"low_p.yaml",
}

var TaskFilenamesMapped = map[string]string{
	"high":   "high_p.yaml",
	"medium": "medium_p.yaml",
	"low":    "low_p.yaml",
}

const fileCount int = len(taskFilenames)

// The name of the YAML file used for storing configs -> TODO
var ConfigFile string = "to_go.cfg.yaml"
var Config *cfg.ConfigValue = new(cfg.ConfigValue)

var TaskList *task.TaskList
