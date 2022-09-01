package cli

import (
	"github.com/leo-alvarenga/to-go/cli/options"
	"github.com/leo-alvarenga/to-go/cli/util"
	"github.com/leo-alvarenga/to-go/engine"
)

/*
Entrypoint to To go's CLI;
Handles most of the decision making and interface calls
based on the args provided by the user.
*/
func CLIEntrypoint(args []string) bool {
	engine.LoadConfig()

	if len(args) > 1 {
		option := args[1]
		modifier := ""

		if len(args) > 2 {
			modifier = args[2]
		}

		if !util.IsThisAnOption(option) {
			return options.InvalidOptionAlert(option)
		}

		engine.StartupEngine()

		switch option {

		case util.List:
			return options.ListOption(modifier)

		default:
			break
		}
	}

	return options.HelpMessage()
}
