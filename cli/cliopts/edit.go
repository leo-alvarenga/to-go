package cliopts

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/ngops"
	"github.com/leo-alvarenga/to-go/shared/styles"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func EditOption() bool {
	ListOption("")
	err := ngops.Edit(getEditInfo())

	if err != nil {
		styles.ShowAsError(ng.Config.Colors, "Unable to edit task", err.Error())
	}

	ListOption("")

	return false
}

func getEditInfo() (n, old task.Task) {
	var choice string

	titles := ng.TaskList.GetAllTitles()
	if len(titles) <= 0 {
		styles.ShowAsError(
			ng.Config.Colors,
			"Hold up, cowboy!",
			"There are no tasks! Add one first if you want to edit them!",
		)

		return
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
		return
	}

	old = *tmp
	n = old

	p, s := getAllPriorities(), getAllStatuses()
	t := old.Title
	var qs = []*survey.Question{
		{
			Name: "Title",
			Prompt: &survey.Input{
				Message: "Type in the title of your task:",
				Default: t,
			},
		},
		{
			Name: "Description",
			Prompt: &survey.Multiline{
				Message: "Type in the description of your task:\n",
				Default: old.Description,
			},
		},
		{
			Name: "Priority",
			Prompt: &survey.Select{
				Message: "Select the priority of your task:",
				Options: p,
				Default: old.Priority,
				Help:    "The value selected will be used to define the task priority.",
			},
		},
		{
			Name: "Status",
			Prompt: &survey.Select{
				Message: "Select the status of your task:",
				Options: s,
				Default: old.Status,
				Help:    "The value selected will be used to define the task status.",
			},
		},
	}

	survey.Ask(qs, &n)

	return
}
