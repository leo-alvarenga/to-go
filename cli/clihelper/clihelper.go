package clihelper

import (
	"fmt"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/shared/styles"
	"github.com/leo-alvarenga/to-go/shared/task"
)

/* Returns a string filled with a finite whitespace (' ') characters based on the 'count' parameter */
func getGapFiller(count int) string {
	s := ""

	for i := 0; i < count; i++ {
		s += " "
	}

	return s
}

/* Breaks a long string in multiple strings of length smaller or equal than 'maxLen' */
func splitLongString(in string, maxLen int) (out []string) {
	l := len(in)

	if l <= maxLen {
		out = append(out, in)
	} else {
		var temp string

		last := 0
		for i := maxLen; i < l; i += maxLen {
			temp = in[last:i]
			last = i

			out = append(out, temp)
		}
	}

	return
}

/*
Displays a task in a styled manner
*/
func DisplayTask(t task.Task) {
	style := new(styles.OutputStyle)
	c := ng.Config
	out := ""

	style.New(c.Colors.Attention, "", []string{"bold"})
	out += style.Style(" > ")

	out += "[" + t.GetPriorityCharacterStyled(c.Colors.Priority) + "] "
	out += "[" + t.GetStatusCharacterStyled(c.Colors.Status, c.UseUnicode) + "] "

	out += style.Style(t.Title) + "\t\t(" + t.CreatedIn + " -"

	if t.FinishedIn == "" {
		out += " [...]"
	} else {
		out += t.FinishedIn
	}

	out += ")"

	fmt.Println(out)
}

/*
Display all the contents of a Task, not leaving out a single character
*/
func DisplayTaskVerbose(t task.Task) {
	fmt.Println("Not implemented yet :(")
}

/* Checks if 's' is a valid CLI option */
func IsThisAnOption(s string) bool {
	for _, opt := range CLIOptions {
		if opt == s {
			return true
		}
	}

	return false
}
