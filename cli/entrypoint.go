package cli

import (
	"fmt"

	"github.com/leo-alvarenga/to-go/api"
	"github.com/leo-alvarenga/to-go/shared"
)

/*
Entrypoint to To go's CLI;
Handles most of the decision making and interface calls
based on the args provided by the user.
*/
func CLIEntrypoint(args []string) bool {

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

func helpMessage() bool {
	shared.ShowLogo()

	return false
}

func invalidOptionAlert(input string) bool {
	fmt.Printf("\"%s\" is not a valid option!\n\n", input)

	return helpMessage()
}
