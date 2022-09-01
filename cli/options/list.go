package options

import (
	"github.com/leo-alvarenga/to-go/cli/util"
	"github.com/leo-alvarenga/to-go/engine"
	"github.com/leo-alvarenga/to-go/shared/task"
)

/* Handles the interface calls for To go's 'list' option */
func ListOption(modifier string) bool {
	switch modifier {
	case util.CLIModifiers["verbose"]:
		return showTasksFull()
	case util.CLIModifiers["header_only"]:
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
	tasks := engine.GetTasks()

	util.DisplayBorder(false)

	util.DisplayTask(task.Task{
		Title:       "Title",
		Description: "Description",
		Priority:    "P...",
		Status:      "Status",
	})

	util.DisplayBorder(true)
	for _, taskList := range tasks {
		for _, todo := range *taskList {
			util.DisplayTask(todo)
		}
	}

	util.DisplayBorder(true)

	return false
}

func showTasksFull() bool {
	tasks := engine.GetTasks()

	util.DisplayBorder(false)

	for _, taskList := range tasks {
		for _, todo := range *taskList {
			util.DisplayTaskVerbose(todo)
		}

		util.DisplayBorder(true)
	}

	return false
}

func showTasksHeaders() bool {
	tasks := engine.GetTasks()

	util.DisplayBorder(false)

	for _, taskList := range tasks {
		for _, todo := range *taskList {
			util.DisplayTaskHeader(todo)
		}
		util.DisplayBorder(true)
	}

	return false
}
