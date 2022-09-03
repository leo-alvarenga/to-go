package ngops

import (
	"errors"
	"fmt"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func Add(t task.Task) error {
	pointers := ng.GetTasks()

	ref := pointers[t.Priority]

	return addTask(ref, t)
}

func addTask(ref *[]task.Task, t task.Task) error {
	for _, item := range *ref {
		if item.Title == t.Title {

			return errors.New(fmt.Sprintf("Task \"%s\" already exists", t.Title))
		}
	}

	id := ""

	if len(*ref) > 0 {
		pos := len(*ref) - 1
		id = (*ref)[pos].Id
	} else {
		id = ng.TaskFilenamesMapped[t.Priority] + "--"
	}

	t.Id = id[:len(id)-1] + string(rune(len(*ref)))
	*ref = append(*ref, t)

	return ng.WriteToYamlFile(ng.TaskFilenamesMapped[t.Priority], ref)
}
