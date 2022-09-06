package clihelper

const (
	Add     string = "add"
	Edit    string = "edit"
	Update  string = "update"
	Remove  string = "remove"
	Finish  string = "finish"
	Help    string = "help"
	List    string = "list"
	Verbose string = "--verbose"
)

var CLIOptions map[string]string = map[string]string{
	Add:    Add,
	Edit:   Edit,
	Update: Update,
	Remove: Remove,
	Finish: Finish,
	Help:   Help,
	List:   List,
}

var CLIModifiers map[string]string = map[string]string{
	"verbose": Verbose,
}
