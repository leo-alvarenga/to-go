package cliopts

import (
	"github.com/leo-alvarenga/to-go/ng"
)

/* Handles the interface calls for To go's 'list' option */
func ListOption() bool {
	return showTasks()
}

/*
Displays all tasks in such a way as to standartize the length of each of its info,
in a table-like output format
*/
func showTasks() bool {
	for _, t := range ng.TaskList.GetAllTasks() {
		t.DisplayTask(*ng.Config)
	}

	return false
}
