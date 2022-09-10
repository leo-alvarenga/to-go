package cliopts

import (
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/ngops"
	"github.com/leo-alvarenga/to-go/shared/styles"
)

func RemoveOption() bool {
	err := ngops.Remove(getRemovalInfo())

	if err != nil {
		styles.ShowAsError(ng.Config.Colors, "Unable to remove task", err.Error())
	}

	return false
}

func getRemovalInfo() string {
	choice, err := selectTask()
	if err != nil {
		styles.ShowAsError(ng.Config.Colors, "Error!", err.Error())
		return ""
	}

	return choice.Title
}
