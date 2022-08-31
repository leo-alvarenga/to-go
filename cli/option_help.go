package cli

import (
	"fmt"

	"github.com/leo-alvarenga/to-go/api"
	"github.com/leo-alvarenga/to-go/shared/styles"
)

func helpMessage() bool {
	colors := api.Config.Colors

	style := new(styles.TextStyle)
	style.New(colors.Success, "", []string{"bold"})

	showLogo(style)

	return false
}

func invalidOptionAlert(input string) bool {
	fmt.Printf("\"%s\" is not a valid option!\n\n", input)

	return helpMessage()
}
