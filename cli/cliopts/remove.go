package cliopts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/ngops"
	"github.com/leo-alvarenga/to-go/shared/styles"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func RemoveOption() bool {
	err := ngops.Remove(getRemovalInfo())

	if err != nil {
		styles.ShowAsError(ng.Config.Colors, "Unable to remove task", err.Error())
	}

	return false
}

func getRemovalInfo() (t task.Task) {
	var choice string

	titles := getAllTitles()
	q1 := []*survey.Question{
		{
			Name: "Target",
			Prompt: &survey.Select{
				Message: "Select the task you want to edit:",
				Options: titles,
				Help:    "The task selected will be the target for any changes you choose to make.",
			},
		},
	}

	survey.Ask(q1, &choice)
	index, priority := getTaskIndex(choice)
	tasks := ng.GetTasks()[priority]

	t = (*tasks)[index]
	return
}
