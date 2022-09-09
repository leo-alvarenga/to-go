package ngops

import (
	"errors"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/storage"
	"github.com/leo-alvarenga/to-go/shared/task"
)

/*
Removes a task from the persistent volume chose by the user.

Returns an error if something in the remove proccess went badly or
if the task does not exists.
*/
func Remove(t task.Task) error {
	if ng.Config.UseSQLite() {
		return storage.RemoveFromSQLite(t)
	}

	var err error
	ng.GetTasks()[t.Priority], err = removeTask(ng.GetTasks()[t.Priority], t)

	if err != nil {
		return err
	}

	return storage.WriteToYamlFile(
		ng.TaskFilenamesMapped[t.Priority],
		ng.GetTasks()[t.Priority],
	)
}

/*
Removes a task 't' from the Task slice that contains it.

  - First return value is an updated copy of the pointer.
  - Second return value signals wheter or not an error has occurred.
*/
func removeTask(ref *[]task.Task, t task.Task) (*[]task.Task, error) {

	index := -1
	for i, item := range *ref {
		if item.Title == t.Title {
			index = i
			break
		}
	}

	if index >= 0 {
		if len(*ref) > 1 {
			t1, t2 := (*ref)[:index-1], (*ref)[index+1:]
			ref = new([]task.Task)

			for _, item := range t1 {
				*ref = append(*ref, item)
			}

			for _, item := range t2 {
				*ref = append(*ref, item)
			}
		} else {
			ref = nil
			ref = new([]task.Task)
		}
	} else {
		return ref, errors.New("Task not found.")
	}

	return ref, nil
}
