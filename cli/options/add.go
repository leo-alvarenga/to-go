package options

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/leo-alvarenga/to-go/engine/operations"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func AddOption() bool {
	err := operations.Add(readInput())

	return err != nil
}

func readInput() (t task.Task) {
	p := []string{}

	for _, priority := range task.Priorities {
		p = append(p, priority)
	}

	var qs = []*survey.Question{
		{
			Name: "Title",
			Prompt: &survey.Input{
				Message: "Type in the title of your task:",
			},
		},
		{
			Name: "Description",
			Prompt: &survey.Multiline{
				Message: "Type in the description of your task:\n",
			},
		},
		{
			Name: "Priority",
			Prompt: &survey.Select{
				Message: "Select the priority of your task:",
				Options: p,
				Help: "The value selected will be used to define the task priority." +
					"Don't worry if you chose the wrong one, you can always edit your task later.",
			},
		},
	}

	survey.Ask(qs, &t)

	return
}
