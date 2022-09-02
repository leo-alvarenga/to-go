package operations

import (
	"errors"
	"fmt"

	"github.com/leo-alvarenga/to-go/engine"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func Add(t task.Task) error {
	var ref *[]task.Task
	pointers := engine.GetTasks()

	switch t.Priority {
	case task.Priorities["high"]:
		ref = pointers[2]
	case task.Priorities["medium"]:
		ref = pointers[1]
	default:
		ref = pointers[0]
	}

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
		id = engine.TaskFilenamesMapped[t.Priority] + "-0"
	}

	t.Id = id[:len(id)-1] + string(rune(len(*ref)))
	*ref = append(*ref, t)

	return engine.WriteToYamlFile(engine.TaskFilenamesMapped[t.Priority], ref)
}
