package cli

import (
	"fmt"

	"github.com/leo-alvarenga/to-go/api"
	"github.com/leo-alvarenga/to-go/shared"
)

func CLIEntrypoint(args []string) bool {

	if len(args) > 1 {
		option := args[1]

		api.StartupEngine()

		switch option {

		default:
			return helpMessage()
		}
	}

	return helpMessage()
}

func helpMessage() bool {
	shared.ShowLogo()

	return false
}

func invalidOptionAlert(input string) bool {
	fmt.Printf("\"%s\" is not a valid option!", input)

	return helpMessage()
}
