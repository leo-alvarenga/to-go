package options

import (
	"github.com/leo-alvarenga/to-go/cli/util"
	"github.com/leo-alvarenga/to-go/engine"
)

/* Handles the interface calls for To go's 'list' option */
func ListOption(modifier string) bool {
	switch modifier {
	case util.CLIModifiers["verbose"]:
		return showTasksVerbose()
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

	for _, t := range *tasks["high"] {
		util.DisplayTask(t)
	}

	for _, t := range *tasks["medium"] {
		util.DisplayTask(t)
	}

	for _, t := range *tasks["low"] {
		util.DisplayTask(t)
	}

	return false
}

func showTasksVerbose() bool {
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
