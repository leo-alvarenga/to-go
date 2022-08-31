package api

import (
	"github.com/leo-alvarenga/to-go/shared"
	"github.com/leo-alvarenga/to-go/shared/config"
)

// The name of the YAML file used for storing tasks with unassigned priority
var taskFiles = [3]string{
	"low_p.yaml",
	"medium_p.yaml",
	"high_p.yaml",
}

const fileCount int = len(taskFiles)

// The name of the YAML file used for storing configs -> TODO
var ConfigFile string = "to_go.cfg.yaml"
var Config *config.ConfigValue = new(config.ConfigValue)

var lowPriorityTasks *[]shared.Task
var mediumPriorityTasks *[]shared.Task
var highPriorityTasks *[]shared.Task
