package ngops

import (
	"time"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/storage"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func Update(t task.Task) error {
	if t.Status == task.Statuses["done"] {
		t.FinishedIn = getDateInToGosFmt(time.Now().Date())
	} else {
		t.FinishedIn = ""
	}

	ng.TaskList.Update(t)
	if ng.Config.UseSQLite() {
		return storage.UpdateOnSQLite(t)
	}

	return writeToYamlWrapper()
}
