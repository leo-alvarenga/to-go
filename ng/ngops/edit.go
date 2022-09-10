package ngops

import (
	"time"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/storage"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func Edit(n, old task.Task) error {
	if n.Status == task.Statuses["done"] {
		n.FinishedIn = getDateInToGosFmt(time.Now().Date())
	}

	if ng.Config.UseSQLite() {
		return storage.EditOnSQLite(n)
	}

	err := ng.TaskList.Edit(old, n)
	if err != nil {
		return nil
	}

	return writeToYamlWrapper()
}
