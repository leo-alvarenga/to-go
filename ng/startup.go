package ng

import (
	"github.com/leo-alvarenga/to-go/ng/storage"
	"github.com/leo-alvarenga/to-go/shared/task"
)

/* Executes all necessary steps to spin up the ng */
func Startup() error {
	TaskList.Low = nil
	TaskList.Medium = nil
	TaskList.High = nil
	TaskList = nil

	TaskList = new(task.TaskList)
	TaskList.New()

	if Config.Storage == "sqlite" {
		return storage.RetriveTasksFromSQLite(TaskList)
	}

	return storage.RetrieveTasksFromYaml(taskFilenames, TaskList)
}

func LoadConfig() error {
	Config.New()
	return Config.LoadFromYaml(ConfigFile)
}
