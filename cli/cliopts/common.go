package cliopts

import (
	"github.com/leo-alvarenga/to-go/shared/task"
)

func getAllPriorities() []string {
	p := []string{}
	for _, i := range task.Priorities {
		p = append(p, i)
	}

	return p
}

func getAllStatuses() []string {
	s := []string{}
	for _, i := range task.Statuses {
		s = append(s, i)
	}

	return s
}
