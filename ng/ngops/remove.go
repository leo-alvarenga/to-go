package ngops

import (
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/storage"
)

/*
Removes a task from the persistent volume chose by the user.

Returns an error if something in the remove proccess went badly or
if the task does not exists.
*/
func Remove(title string) error {
	if ng.Config.UseSQLite() {
		t, err := ng.TaskList.GetTaskByTitle(title)
		if err != nil {
			return err
		}

		return storage.RemoveFromSQLite(t)
	}

	err := ng.TaskList.Remove(title)
	if err != nil {
		return err
	}

	return writeToYamlWrapper()
}
