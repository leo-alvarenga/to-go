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
	}

	if ng.Config.UseSQLite() {
		return storage.UpdateOnSQLite(t)
	}

	tasks := ng.GetTasks()[t.Priority]

	for i, ts := range *tasks {
		if ts.Id == t.Id {
			(*tasks)[i].Status = t.Status
		}
	}

	return storage.WriteToYamlFile(ng.TaskFilenamesMapped[t.Priority], tasks)
}
