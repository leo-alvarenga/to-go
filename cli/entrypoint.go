package cli

import (
	"github.com/leo-alvarenga/to-go/api"
)

/*
Entrypoint to To go's CLI;
Handles most of the decision making and interface calls
based on the args provided by the user.
*/
func CLIEntrypoint(args []string) bool {
	api.LoadConfig()

	if len(args) > 1 {
		option := args[1]
		modifier := ""

		if len(args) > 2 {
			modifier = args[2]
		}

		if !isThisAnOption(option) {
			return invalidOptionAlert(option)
		}

		api.StartupEngine()

		switch option {

		case list:
			return listOption(modifier)

		default:
			break
		}
	}

	return helpMessage()
}
