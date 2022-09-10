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
			ng.Config.Colors,
			"An error has occurred",
			err.Error(),
		)
	}

	if !term.IsTerminal(int(syscall.Stdin)) {
		styles.ShowAsError(
			ng.Config.Colors,
			"It seems you are not in interactive mode",
			"To use To go, you need to be use your terminal/prompt in interactive mode",
		)
	}

	if len(args) > 1 {
		option := args[1]

		if !clihelper.IsThisAnOption(option) {
			return cliopts.InvalidOptionAlert(option)
		}

		err = ng.Startup()

		if err != nil {
			styles.ShowAsError(ng.Config.Colors, "An error has occurred", err.Error())

			return true
		}

		switch option {
		case clihelper.List:
			return cliopts.ListOption()

		case clihelper.Describe:
			return cliopts.DescribeOption()

		case clihelper.Add:
			return cliopts.AddOption()

		case clihelper.Edit:
			return cliopts.EditOption()

		case clihelper.Update:
			return cliopts.UpdateOption()

		case clihelper.Finish:
			return cliopts.FinishOption()

		case clihelper.Remove:
			return cliopts.RemoveOption()

		default:
			break
		}
	}

	return cliopts.HelpMessage()
}
