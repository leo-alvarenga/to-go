package api

import (
	"github.com/leo-alvarenga/to-go/shared"
)

/* Executes all necessary steps to spin up the engine */
func StartupEngine() {
	lowPriorityTasks = new([]shared.Task)
	mediumPriorityTasks = new([]shared.Task)
	highPriorityTasks = new([]shared.Task)

	retrieveTasks()
}

func LoadConfig() {
	Config.New()
	Config.LoadFromYaml(ConfigFile)
}
