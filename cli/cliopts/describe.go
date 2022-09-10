package cliopts

import (
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/shared/styles"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func DescribeOption() bool {
	getTaskInfoToDescribe().Describe(*ng.Config)

	return false
}

func getTaskInfoToDescribe() task.Task {
	choice, err := selectTask()
	if err != nil {
		styles.ShowAsError(ng.Config.Colors, "Error!", err.Error())
	}

	return choice
}
