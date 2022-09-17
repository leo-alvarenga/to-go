package cliopts

import (
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/ngops"
)

func DashboardOption() bool {
	ngops.Dashboard(ng.Config)

	return false
}
