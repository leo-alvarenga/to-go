package clihelper

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

func IsThisAnOption(opt string) bool {
	for _, o := range CLIOptions {
		if o == opt {
			return true
		}
	}

	return false
}
