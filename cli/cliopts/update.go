package cliopts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/ngops"
	"github.com/leo-alvarenga/to-go/shared/styles"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func UpdateOption() bool {
	err := ngops.Update(getUpdateInfo())

	if err != nil {
		styles.ShowAsError(ng.Config.Colors, "Unable to update task status", err.Error())
	}

	return false
}

func getUpdateInfo() (t task.Task) {
	t, err := selectTask()
	if err != nil {
		styles.ShowAsError(ng.Config.Colors, "Error!", err.Error())
		return
	}

	s := getAllStatuses()
	var qs = []*survey.Question{
		{
			Name: "Status",
			Prompt: &survey.Select{
				Message: "Select the status of your task:",
				Options: s,
				Default: t.Status,
				Help:    "The value selected will be used to define the task status.",
			},
		},
	}

	survey.Ask(qs, &t)
	return
}
