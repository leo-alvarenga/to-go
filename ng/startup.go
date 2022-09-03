package ng

import (
	"github.com/leo-alvarenga/to-go/shared/task"
)

/* Executes all necessary steps to spin up the ng */
func Startup() error {
	lowPriorityTasks = new([]task.Task)
	mediumPriorityTasks = new([]task.Task)
	highPriorityTasks = new([]task.Task)

	return retrieveTasks()
}

func LoadConfig() error {
	Config.New()
	return Config.LoadFromYaml(ConfigFile)
}
