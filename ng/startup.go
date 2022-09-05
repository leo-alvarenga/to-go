package ng

import (
	"github.com/leo-alvarenga/to-go/ng/storage"
	"github.com/leo-alvarenga/to-go/shared/task"
)

/* Executes all necessary steps to spin up the ng */
func Startup() error {
	lowPriorityTasks = new([]task.Task)
	mediumPriorityTasks = new([]task.Task)
	highPriorityTasks = new([]task.Task)

	if Config.Storage == "sqlite" {
		storage.CreateDB()
		return nil
	}

	return retrieveTasksFromYaml()
}

func LoadConfig() error {
	Config.New()
	return Config.LoadFromYaml(ConfigFile)
}
