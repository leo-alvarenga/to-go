package util

import (
	"fmt"

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
Display an input string in such a way as to force said string to not disrespect the
To go's table-like output format
*/
func displayFormatted(format string) {
	var fill string

	var f []string
	l := len(format)

	if l >= maxContentLength {

		var count int

		for i := 0; i < l; i += maxContentLength {
			count = i + maxContentLength
			if count >= l {
				count = count - (count - l)
			}

			f = append(f, format[i:count])
		}
	} else {
		f = append(f, format)
	}

	for _, line := range f {
		fill = ""

		if len(line) < maxContentLength {
			for i := 0; i < (maxContentLength - len(line)); i++ {
				fill += " "
			}
		}

		fmt.Println("| " + line + fill + " |")
	}
}

/*
Displays a task in such a way as to standartize the length of each of its info, while respecting
To go's table-like output format
*/
func DisplayTask(t task.Task) {
	separator := " | "
	maxLenTitle := (maxContentLength -
		(task.LenLongestPriority + task.LenLongestStatus + 9)) / 2

	maxLenDesc := maxLenTitle

	if len(t.Title) > maxLenTitle {
		t.Title = t.Title[:(maxLenTitle-3)] + "..."
	} else if len(t.Title) < maxLenTitle {
		t.Title +=
			getGapFiller(maxLenTitle - len(t.Title))
	}

	if len(t.Description) > maxLenDesc {
		t.Description = t.Description[:maxLenDesc-3] + "..."
	} else if len(t.Description) < maxLenDesc {
		t.Description +=
			getGapFiller(maxLenDesc - len(t.Description))
	}

	if len(t.Priority) < task.LenLongestPriority {
		t.Priority +=
			getGapFiller(task.LenLongestPriority - len(t.Priority))
	}

	if len(t.Status) < task.LenLongestStatus {
		t.Status +=
			getGapFiller(task.LenLongestStatus - len(t.Status))
	}

	out := t.Title +
		separator + t.Description +
		separator + t.Priority +
		separator + t.Status

	displayFormatted(out)
}

/*
Display all the contents of a Task, not leaving out a single character
*/
func DisplayTaskVerbose(t task.Task) {
	separator := "-------------"

	s := splitLongString(t.Id, maxContentLength)
	s = append(s, separator)

	for _, line := range s {
		displayFormatted(line)
	}

	s = splitLongString(t.Title, maxContentLength)
	s = append(s, separator)

	for _, line := range s {
		displayFormatted(line)
	}

	displayFormatted(t.Priority + " -> " + t.Status)
	displayFormatted(separator)

	s = splitLongString(t.Description, maxContentLength)
	s = append(s, separator)

	for _, line := range s {
		displayFormatted(line)
	}
}

/*
Display header contents of a Task, such as Id, Title and priority
*/
func DisplayTaskHeader(t task.Task) {
	separator := " - "
	idLen := len(t.Id)
	out := ""

	l := maxContentLength - (task.LenLongestPriority + len(separator))

	if idLen < l {
		out += t.Id + separator
		remaining := l - len(out)

		if len(t.Title) > remaining {
			out += t.Title[:remaining-len(separator)] + "..."
		} else {
			out += t.Title
		}

		out += separator + t.Priority
	}

	displayFormatted(out)
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
