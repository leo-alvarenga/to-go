package engine

import (
	"github.com/leo-alvarenga/to-go/shared/task"
)

/* Executes all necessary steps to spin up the engine */
func Startup() {
	lowPriorityTasks = new([]task.Task)
	mediumPriorityTasks = new([]task.Task)
	highPriorityTasks = new([]task.Task)

	retrieveTasks()
}

func LoadConfig() {
	Config.New()
	Config.LoadFromYaml(ConfigFile)
}
