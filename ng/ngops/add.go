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
	y, m, d := time.Now().Date()
	extra := ""

	switch d % 10 {
	case 1:
		extra += "st"
	case 2:
		extra += "nd"
	case 3:
		extra += "rd"
	default:
		extra += "th"
	}

	t.Status = task.Statuses["pending"]
	t.CreatedIn = fmt.Sprintf("%s %d%s, %d", m.String(), d, extra, y)
	t.FinishedIn = ""

	for _, item := range *ref {
		if item.Title == t.Title {
			return errors.New(fmt.Sprintf("Task \"%s\" already exists", t.Title))
		}
	}

	if useSQLite {
		return storage.WriteToSQLite(t)
	}

	*ref = append(*ref, t)
	return storage.WriteToYamlFile(ng.TaskFilenamesMapped[t.Priority], ref)
}
