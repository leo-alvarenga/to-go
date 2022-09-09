package ngops

import (
	"time"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/storage"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func Add(t task.Task) error {
	t.Status = task.Statuses["pending"]
	t.CreatedIn = getDateInToGosFmt(time.Now().Date())
	t.FinishedIn = ""

	err := ng.TaskList.Add(t)

	if err != nil {
		return err
	}

	if ng.Config.UseSQLite() {
		return storage.WriteToSQLite(t)
	}

	return writeToYamlWrapper()
}
