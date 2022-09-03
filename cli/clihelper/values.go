package clihelper

import "github.com/leo-alvarenga/to-go/shared/cfg"

var maxLineLength int = cfg.DefaultMaxLineLen // must be an odd number!
var maxContentLength int = maxLineLength - 4

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
