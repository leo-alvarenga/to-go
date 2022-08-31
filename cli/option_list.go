package cli

import (
	"github.com/leo-alvarenga/to-go/api"
	"github.com/leo-alvarenga/to-go/shared"
)

/* Handles the interface calls for To go's 'list' option */
func listOption(modifier string) bool {
	switch modifier {
	case verbose:
		return showTasksFull()
	case header_only:
		return showTasksHeaders()
	default:
		return showTasks()
	}
}

/*
Displays all tasks in such a way as to standartize the length of each of its info,
in a table-like output format
*/
func showTasks() bool {
	tasks := api.GetTasks()

	displayBorder(false)

	displayTask(shared.Task{
		Title:       "Title",
		Description: "Description",
		Priority:    "P...",
		Status:      "Status",
	})

	displayBorder(true)
	for _, taskList := range tasks {
		for _, todo := range *taskList {
			displayTask(todo)
		}
	}

	displayBorder(true)

	return false
}

func showTasksFull() bool {
	tasks := api.GetTasks()

	displayBorder(false)

	for _, taskList := range tasks {
		for _, todo := range *taskList {
			displayTaskVerbose(todo)
		}
		displayBorder(true)
	}

	return false
}

func showTasksHeaders() bool {
	tasks := api.GetTasks()

	displayBorder(false)

	for _, taskList := range tasks {
		for _, todo := range *taskList {
			displayTaskHeader(todo)
		}
		displayBorder(true)
	}

	return false
}
