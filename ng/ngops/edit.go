package ngops

import (
	"errors"
	"fmt"
	"time"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/storage"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func Edit(old, n task.Task) error {
	if n.Status == task.Statuses["done"] {
		n.FinishedIn = getDateInToGosFmt(time.Now().Date())
	}

	if ng.Config.UseSQLite() {
		return storage.EditOnSQLite(n)
	}

	tasks := ng.GetTasks()[old.Priority]
	index := -1
	for i, item := range *tasks {
		fmt.Println(item.Title, old.Title)
		if item.Title == old.Title {
			index = i
			break
		}
	}

	if old.Priority == n.Priority {
		if index >= 0 {
			(*tasks)[index] = n

			return storage.WriteToYamlFile(ng.TaskFilenamesMapped[n.Priority], tasks)
		}
	} else {
		if index >= 0 {
			t1, t2 := (*tasks)[:index-1], (*tasks)[index+1:]
			tasks = new([]task.Task)

			for _, item := range t1 {
				*tasks = append(*tasks, item)
			}

			for _, item := range t2 {
				*tasks = append(*tasks, item)
			}

			storage.WriteToYamlFile(ng.TaskFilenamesMapped[old.Priority], tasks)

			tasks = ng.GetTasks()[n.Priority]
			*tasks = append(*tasks, n)
			return storage.WriteToYamlFile(ng.TaskFilenamesMapped[n.Priority], tasks)
		}
	}

	return errors.New("Task does not exist. Please, try again.")
}
