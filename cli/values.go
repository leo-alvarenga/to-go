package cli

/* The max length a line can have when being printed to stdout */
const defaultMaxLineLen int = 59

var maxLineLength int = defaultMaxLineLen // must be an odd number!
var maxContentLength int = maxLineLength - 4

const (
	add         string = "add"
	delete      string = "delete"
	finish      string = "finish"
	help        string = "help"
	list        string = "list"
	verbose     string = "--verbose"
	header_only string = "--header"
)

var CLIOptions map[string]string = map[string]string{
	add:    add,
	delete: delete,
	finish: finish,
	help:   help,
	list:   list,
}

var CLIModifiers map[string][]string = map[string][]string{
	list: {verbose, header_only},
}
