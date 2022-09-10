package cliopts

import (
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
	styles.ShowAsError(ng.Config.Colors, "\""+input+"\" is not a valid option", "")

	return HelpMessage()
}
