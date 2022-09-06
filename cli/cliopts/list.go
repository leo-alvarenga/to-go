package cliopts

import (
	"github.com/leo-alvarenga/to-go/cli/clihelper"
	"github.com/leo-alvarenga/to-go/ng"
)

/* Handles the interface calls for To go's 'list' option */
func ListOption(modifier string) bool {
	switch modifier {
	case clihelper.CLIModifiers["verbose"]:
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
	tasks := ng.GetTasks()

	for _, t := range *tasks["high"] {
		clihelper.DisplayTask(t)
	}

	for _, t := range *tasks["medium"] {
		clihelper.DisplayTask(t)
	}

	for _, t := range *tasks["low"] {
		clihelper.DisplayTask(t)
	}

	return false
}

func showTasksVerbose() bool {
	tasks := ng.GetTasks()

	for _, taskList := range tasks {
		for _, todo := range *taskList {
			clihelper.DisplayTaskVerbose(todo)
		}
	}

	return false
}
