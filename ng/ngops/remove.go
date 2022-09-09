package ngops

import (
	"fmt"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/storage"
)

/*
Removes a task from the persistent volume chose by the user.

Returns an error if something in the remove proccess went badly or
if the task does not exists.
*/
func Remove(title string) error {
	t, err := ng.TaskList.GetTaskByTitle(title)
	if err != nil {
		return err
	}

	if ng.Config.UseSQLite() {
		return storage.RemoveFromSQLite(*t)
	}

	fmt.Println(*&t.Title)
	err = ng.TaskList.Remove(*t)
	if err != nil {
		return err
	}

	return writeToYamlWrapper()
}
