package ng

import (
	"github.com/leo-alvarenga/to-go/ng/storage"
	"github.com/leo-alvarenga/to-go/shared/task"
)

/*
Returns pointers for each one of the dynamically allocated Task slices
*/
func GetTasks() map[string]*[]task.Task {
	return map[string]*[]task.Task{
		"high":   highPriorityTasks,
		"medium": mediumPriorityTasks,
		"low":    lowPriorityTasks,
	}
}

/* Executes all necessary steps to spin up the ng */
func Startup() error {
	lowPriorityTasks = new([]task.Task)
	mediumPriorityTasks = new([]task.Task)
	highPriorityTasks = new([]task.Task)

	if Config.Storage == "sqlite" {
		return storage.RetriveTasksFromSQLite(
			highPriorityTasks,
			mediumPriorityTasks,
			lowPriorityTasks,
		)
	}

	return storage.RetrieveTasksFromYaml(
		taskFilenames,
		highPriorityTasks,
		mediumPriorityTasks,
		lowPriorityTasks,
	)
}

func LoadConfig() error {
	Config.New()
	return Config.LoadFromYaml(ConfigFile)
}
