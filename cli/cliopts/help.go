package cliopts

import (
	"fmt"

	"github.com/leo-alvarenga/to-go/cli/clihelper"
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/shared/styles"
)

func HelpMessage() bool {
	colors := ng.Config.Colors

	style := new(styles.OutputStyle)
	style.New(colors.Success, "", []string{"bold"})

	clihelper.ShowLogo(style)

	return false
}

func InvalidOptionAlert(input string) bool {
	fmt.Printf("\"%s\" is not a valid option!\n\n", input)

	return HelpMessage()
}
