package cliopts

import (
	"fmt"
	"time"

	"github.com/leo-alvarenga/to-go/cli/clihelper"
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/shared/styles"
)

func HelpMessage() bool {
	colors := ng.Config.Colors

	style := new(styles.OutputStyle)
	alt := new(styles.OutputStyle)

	style.New(colors.Attention, "", []string{"bold"})
	alt.New("green", "", []string{})

	clihelper.ShowLogo(style)
	time.Sleep(1 * time.Second)

	style.New(colors.Success, "", []string{"bold", "underline"})
	style.ShowWithStyle("\n\nOptions:")

	for label, opt := range clihelper.CLIOptionsDescription {
		style.ShowWithStyle(label)

		fmt.Println("\t" + opt[0])

		alt.ShowWithStyle("\t>> " + opt[1])
	}

	return false
}

func InvalidOptionAlert(input string) bool {
	styles.ShowAsError(ng.Config.Colors, "\""+input+"\" is not a valid option", "")

	return HelpMessage()
}
