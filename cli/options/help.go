package options

import (
	"fmt"

	"github.com/leo-alvarenga/to-go/engine"
	"github.com/leo-alvarenga/to-go/shared/logo"
	"github.com/leo-alvarenga/to-go/shared/styles"
)

func HelpMessage() bool {
	colors := engine.Config.Colors

	style := new(styles.OutputStyle)
	style.New(colors.Success, "", []string{"bold"})

	logo.ShowLogo(style)

	return false
}

func InvalidOptionAlert(input string) bool {
	fmt.Printf("\"%s\" is not a valid option!\n\n", input)

	return HelpMessage()
}
