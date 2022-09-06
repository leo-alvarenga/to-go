package cliopts

import (
	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func getAllTitles() []string {
	tasks := ng.GetTasks()
	out := []string{}

	for _, slice := range tasks {
		for _, t := range *slice {
			out = append(out, t.Title)
		}
	}

	return out
}

func getTaskIndex(title string) (int, string) {
	tasks := ng.GetTasks()

	for _, p := range tasks {
		for i, t := range *p {
			if t.Title == title {
				return i, t.Priority
			}
		}
	}

	return -1, ""
}

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
