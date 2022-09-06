package ngops

import (
	"errors"
	"fmt"
	"time"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/storage"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func Add(t task.Task) error {
	pointers := ng.GetTasks()

	ref := pointers[t.Priority]

	return addTask(ref, t, ng.Config.UseSQLite())
}

func addTask(ref *[]task.Task, t task.Task, useSQLite bool) error {
	t.Status = task.Statuses["pending"]
	t.CreatedIn = getDateInToGosFmt(time.Now().Date())
	t.FinishedIn = ""

	for _, item := range *ref {
		if item.Title == t.Title {
			return errors.New(fmt.Sprintf("Task \"%s\" already exists", t.Title))
		}
	}

	if useSQLite {
		return storage.WriteToSQLite(t)
	}

	t.Id = len(*ref)
	*ref = append(*ref, t)
	return storage.WriteToYamlFile(ng.TaskFilenamesMapped[t.Priority], ref)
}
