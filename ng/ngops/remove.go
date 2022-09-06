package ngops

import (
	"errors"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/storage"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func Remove(t task.Task) error {
	if ng.Config.UseSQLite() {
		return storage.RemoveFromSQLite(t)
	}

	tasks := ng.GetTasks()[t.Priority]

	index := -1
	for i, item := range *tasks {
		if item.Title == t.Title {
			index = i
			break
		}
	}

	if index >= 0 {
		if len(*tasks) > 1 {
			t1, t2 := (*tasks)[:index-1], (*tasks)[index+1:]
			tasks = new([]task.Task)

			for _, item := range t1 {
				*tasks = append(*tasks, item)
			}

			for _, item := range t2 {
				*tasks = append(*tasks, item)
			}
		} else {
			tasks = new([]task.Task)
		}
	} else {
		return errors.New("Task not found.")
	}

	return storage.WriteToYamlFile(
		ng.TaskFilenamesMapped[t.Priority],
		tasks,
	)
}
