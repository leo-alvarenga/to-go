package clihelper

const (
	Add     string = "add"
	Delete  string = "delete"
	Finish  string = "finish"
	Help    string = "help"
	List    string = "list"
	Verbose string = "--verbose"
)

var CLIOptions map[string]string = map[string]string{
	Add:    Add,
	Delete: Delete,
	Finish: Finish,
	Help:   Help,
	List:   List,
}

var CLIModifiers map[string]string = map[string]string{
	"verbose": Verbose,
}
