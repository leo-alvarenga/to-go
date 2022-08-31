package styles

func generateANSISequence(txt, bg string, style []string) (string, string) {

	t, b, s := "", "", ""

	for key := range ansiColors {
		if txt == key {
			t = ansiColors["txt"] + ansiColors[key] + ansiColors["separator"]
		}

		if bg == key {
			b = ansiColors["bg"] + ansiColors[key] + ansiColors["separator"]
		}
	}

	for key := range ansiColors {
		for _, st := range style {
			if st == key {
				s += ansiColors[key] + ansiColors["separator"]
			}
		}
	}

	if len(s) > 0 {
		s = s[:len(s)-1]
	} else if len(b) > 0 {
		b = b[:len(b)-1]
	} else if len(t) > 0 {
		t = t[:len(t)-1]
	} else {
		return "", "\033[0m"
	}

	return ansiColors["escape"] + t + b + s + ansiColors["end"], "\033[0m"
}
