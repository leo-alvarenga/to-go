package clihelper

import "fmt"

const (
	Executable string = "to-go"
	Help       string = "help"
	Add        string = "add"
	Remove     string = "remove"
	Edit       string = "edit"
	Update     string = "update"
	Finish     string = "finish"
	List       string = "list"
	Describe   string = "describe"
	Dashboard  string = "dashboard"
)

var CLIOptions map[string]string = map[string]string{
	"help":      Help,
	"add":       Add,
	"remove":    Remove,
	"edit":      Edit,
	"update":    Update,
	"finish":    Finish,
	"list":      List,
	"describe":  Describe,
	"dashboard": Dashboard,
}

var CLIOptionsDescription map[string][2]string = map[string][2]string{
	"help": {
		"Shows helpful information on the options available",
		fmt.Sprintf("%s %s", Executable, Help),
	},
	"add": {
		"Adds a task based of on the users input",
		fmt.Sprintf("%s %s", Executable, Add),
	},
	"remove": {
		"Removes a task selected by the user",
		fmt.Sprintf("%s %s", Executable, Remove),
	},
	"edit": {
		"Edits a task selected by the user",
		fmt.Sprintf("%s %s", Executable, Edit),
	},
	"update": {
		"Updates the status of a task selected by the user",
		fmt.Sprintf("%s %s", Executable, Update),
	},
	"finish": {
		"Updates the status of a task selected by the user to 'done'",
		fmt.Sprintf("%s %s", Executable, Finish),
	},
	"list": {
		"Lists all the tasks, displaying their titles, priorities, statuses and dates",
		fmt.Sprintf("%s %s", Executable, List),
	},
	"describe": {
		"Displays all the info pertaining to a task chosen by you",
		fmt.Sprintf("%s %s", Executable, Describe),
	},
}
