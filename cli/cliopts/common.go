package cliopts

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/shared/styles"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func getAllPriorities() []string {
	p := []string{}
	for _, i := range task.Priorities {
		p = append(p, i)
	}

	return p
}

func getAllStatuses() []string {
	s := []string{}
	for _, i := range task.Statuses {
		s = append(s, i)
	}

	return s
}

func selectTask() (task.Task, error) {
	var choice string

	titles := ng.TaskList.GetAllTitles()
	if len(titles) <= 0 {
		styles.ShowAsError(
			ng.Config.Colors,
			"Hold up, cowboy!",
			"There are no tasks! Add one first if you want to do something!",
		)

		return task.Task{}, errors.New("No tasks available")
	}

	q1 := []*survey.Question{
		{
			Name: "Target",
			Prompt: &survey.Select{
				Message: "Select a task:",
				Options: titles,
				Help:    "The task selected will be the target on this operation.",
			},
		},
	}

	survey.Ask(q1, &choice)

	t, err := ng.TaskList.GetTaskByTitle(choice)
	if err != nil {
		return task.Task{}, err
	}

	return t, nil
}
