package cliopts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/ngops"
	"github.com/leo-alvarenga/to-go/shared/styles"
)

func RemoveOption() bool {
	err := ngops.Remove(getRemovalInfo())

	if err != nil {
		styles.ShowAsError(ng.Config.Colors, "Unable to remove task", err.Error())
	}

	return false
}

func getRemovalInfo() string {
	var choice string

	titles := ng.TaskList.GetAllTitles()
	if len(titles) <= 0 {
		styles.ShowAsError(
			ng.Config.Colors,
			"Hold up, cowboy!",
			"There are no tasks! Add one first if you want to remove one!",
		)

		return ""
	}

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

	tmp, err := ng.TaskList.GetTaskByTitle(choice)
	if err != nil {
		styles.ShowAsError(ng.Config.Colors, "Error!", err.Error())
		return ""
	}

	return tmp.Title
}
