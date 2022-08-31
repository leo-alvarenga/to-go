package main

import (
	"os"

	"github.com/leo-alvarenga/to-go/cli"
	"github.com/leo-alvarenga/to-go/shared/config"
)

func main() {
	var config *config.ConfigValue
	cli.CLIEntrypoint(os.Args, config)
}
