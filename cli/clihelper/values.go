package clihelper

import "fmt"

const (
	Executable string = "to-go"
	Add        string = "add"
	Edit       string = "edit"
	Update     string = "update"
	Remove     string = "remove"
	Finish     string = "finish"
	Help       string = "help"
	List       string = "list"
	Describe   string = "describe"
)

var CLIOptions map[string]string = map[string]string{
	"add":      Add,
	"edit":     Edit,
	"update":   Update,
	"remove":   Remove,
	"finish":   Finish,
	"help":     Help,
	"list":     List,
	"describe": Describe,
}

var CLIOptionsDescription map[string][2]string = map[string][2]string{
	"help": {
		"Adds a task based of on the users input",
		fmt.Sprintf("%s %s", Executable, Help),
	},
	"add": {
		"Adds a task based of on the users input",
		fmt.Sprintf("%s %s", Executable, Add),
	},
	"remove": {
		"Adds a task based of on the users input",
		fmt.Sprintf("%s %s", Executable, Remove),
	},
	"edit": {
		"Edits a task chosen by you",
		fmt.Sprintf("%s %s", Executable, Edit),
	},
	"update": {
		"Updates the status of a task chosen by you",
		fmt.Sprintf("%s %s", Executable, Update),
	},
	"finish": {
		"Updates the status of a task chosen by you to 'done'",
		fmt.Sprintf("%s %s", Executable, Finish),
	},
	"list": {
		"Lists all the tasks, including ther titles, priorities, statuses and dates",
		fmt.Sprintf("%s %s", Executable, List),
	},
	"describe": {
		"Displays all the info pertaining to a task chosen by you",
		fmt.Sprintf("%s %s", Executable, Describe),
	},
}
