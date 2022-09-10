package cliopts

import (
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/ngops"
	"github.com/leo-alvarenga/to-go/shared/styles"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func FinishOption() bool {
	choice, err := selectTask()
	if err != nil {
		styles.ShowAsError(ng.Config.Colors, "Unable to update task status", err.Error())
		return true
	}

	choice.Status = task.Statuses["done"]
	err = ngops.Update(choice)

	if err != nil {
		styles.ShowAsError(ng.Config.Colors, "Unable to update task status", err.Error())
		return true
	}

	return false
}
