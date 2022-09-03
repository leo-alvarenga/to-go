package util

import (
	"fmt"

	"github.com/leo-alvarenga/to-go/engine"
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
Display a crude border-like character sequence;

The 'end' value indicates whether or not the border corresponds to
the bottom
*/
func DisplayBorder(end bool) {
	out := ""
	limit := ""

	if end {
		limit = "|"
	} else {
		limit = " "
	}

	out += limit
	for i := 1; i < maxLineLength-1; i++ {
		out += "_"
	}

	out += limit
	fmt.Println(out)
}

/*
Displays a task in a styled manner
*/
func DisplayTask(t task.Task) {
	style := new(styles.OutputStyle)
	colors := engine.Config.Colors
	out := ""

	style.New(colors.Accent, "", []string{"bold"})
	out += style.Style(" > ")

	out += "[" + t.GetPriorityCharacterStyled() + "] "
	out += "[" + t.GetStatusCharacterStyled() + "] "

	out += t.Title

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
