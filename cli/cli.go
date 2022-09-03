package cli

import (
	"syscall"

	"github.com/leo-alvarenga/to-go/cli/clihelper"
	"github.com/leo-alvarenga/to-go/cli/cliopts"
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/shared/styles"
	"golang.org/x/term"
)

/*
Entrypoint to To go's CLI;
Handles most of the decision making and interface calls
based on the args provided by the user.
*/
func Entrypoint(args []string) bool {
	err := ng.LoadConfig()

	if err != nil {
		styles.ShowAsError(
			"An error has occurred",
			err.Error(),
		)
	}

	if !term.IsTerminal(int(syscall.Stdin)) {
		styles.ShowAsError(
			"It seems you are not in interactive mode",
			"To use To go, you need to be use your terminal/prompt in interactive mode",
		)
	}

	if len(args) > 1 {
		option := args[1]
		modifier := ""

		if len(args) > 2 {
			modifier = args[2]
		}

		if !clihelper.IsThisAnOption(option) {
			return cliopts.InvalidOptionAlert(option)
		}

		err = ng.Startup()

		if err != nil {
			styles.ShowAsError(
				"An error has occurred",
				err.Error(),
			)

			return true
		}

		switch option {
		case clihelper.List:
			return cliopts.ListOption(modifier)

		case clihelper.Add:
			return cliopts.AddOption()

		default:
			break
		}
	}

	return cliopts.HelpMessage()
}
