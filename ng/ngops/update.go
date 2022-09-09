package ngops

import (
	"errors"
	"time"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/storage"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func Update(t task.Task) error {
	if t.Status == task.Statuses["done"] {
		t.FinishedIn = getDateInToGosFmt(time.Now().Date())
	}

	if ng.Config.UseSQLite() {
		return storage.UpdateOnSQLite(t)
	}

	tasks := ng.GetTasks()[t.Priority]

	err := updateTask(tasks, t)

	if err != nil {
		return err
	}

	return storage.WriteToYamlFile(ng.TaskFilenamesMapped[t.Priority], tasks)
}

func updateTask(ref *[]task.Task, t task.Task) error {
	for i, ts := range *ref {
		if ts.Id == t.Id {
			(*ref)[i].Status = t.Status
			return nil
		}
	}

	return errors.New("Task not found.")
}
