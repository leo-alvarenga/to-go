package cli

import (
	"fmt"

	"github.com/leo-alvarenga/to-go/shared"
	"github.com/leo-alvarenga/to-go/shared/styles"
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
func displayBorder(end bool) {
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
func displayTask(task shared.Task) {
	separator := " | "
	maxLenTitle := (maxContentLength -
		(shared.LenLongestPriority + shared.LenLongestStatus + 9)) / 2

	maxLenDesc := maxLenTitle

	if len(task.Title) > maxLenTitle {
		task.Title = task.Title[:(maxLenTitle-3)] + "..."
	} else if len(task.Title) < maxLenTitle {
		task.Title +=
			getGapFiller(maxLenTitle - len(task.Title))
	}

	if len(task.Description) > maxLenDesc {
		task.Description = task.Description[:maxLenDesc-3] + "..."
	} else if len(task.Description) < maxLenDesc {
		task.Description +=
			getGapFiller(maxLenDesc - len(task.Description))
	}

	if len(task.Priority) < shared.LenLongestPriority {
		task.Priority +=
			getGapFiller(shared.LenLongestPriority - len(task.Priority))
	}

	if len(task.Status) < shared.LenLongestStatus {
		task.Status +=
			getGapFiller(shared.LenLongestStatus - len(task.Status))
	}

	out := task.Title +
		separator + task.Description +
		separator + task.Priority +
		separator + task.Status

	displayFormatted(out)
}

/*
Display all the contents of a Task, not leaving out a single character
*/
func displayTaskVerbose(task shared.Task) {
	separator := "-------------"

	s := splitLongString(task.Id, maxContentLength)
	s = append(s, separator)

	for _, line := range s {
		displayFormatted(line)
	}

	s = splitLongString(task.Title, maxContentLength)
	s = append(s, separator)

	for _, line := range s {
		displayFormatted(line)
	}

	displayFormatted(task.Priority + " -> " + task.Status)
	displayFormatted(separator)

	s = splitLongString(task.Description, maxContentLength)
	s = append(s, separator)

	for _, line := range s {
		displayFormatted(line)
	}
}

/*
Display header contents of a Task, such as Id, Title and priority
*/
func displayTaskHeader(task shared.Task) {
	separator := " - "
	idLen := len(task.Id)
	out := ""

	l := maxContentLength - (shared.LenLongestPriority + len(separator))

	if idLen < l {
		out += task.Id + separator
		remaining := l - len(out)

		if len(task.Title) > remaining {
			out += task.Title[:remaining-len(separator)] + "..."
		} else {
			out += task.Title
		}

		out += separator + task.Priority
	}

	displayFormatted(out)
}

/*
Checks if 'mod' is a valid modifier for the 'option'
-> TODO: Prob should remove this, not being used...
*/
func isThisAModifierForThis(mod, option string) bool {
	for _, m := range CLIModifiers[option] {
		if m == mod {
			return true
		}
	}

	return false
}

/* Checks if 's' is a valid CLI option */
func isThisAnOption(s string) bool {
	for _, opt := range CLIOptions {
		if opt == s {
			return true
		}
	}

	return false
}

func showWithStyle(s string, style *styles.TextStyle) {
	fmt.Print(style.ANSI + s)
	fmt.Println(style.Reset)
}
