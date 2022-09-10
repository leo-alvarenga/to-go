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
	var choice string

	titles := ng.TaskList.GetAllTitles()
	if len(titles) <= 0 {
		st := new(styles.OutputStyle)

		st.New("red", "", []string{"bold", "underline"})
		st.ShowWithStyle("Hold up, cowboy!")

		st.New("red", "", []string{"bold"})
		st.ShowWithStyle("There are no tasks! Add one first if you want to update one!")

		st = nil
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

	t, err := ng.TaskList.GetTaskByTitle(choice)
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
