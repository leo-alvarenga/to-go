package styles

/* Values used for ANSI escape sequence generation */
var ansiColors = map[string]string{
	"black":   "0",
	"red":     "1",
	"green":   "2",
	"yellow":  "3",
	"blue":    "4",
	"purple":  "5",
	"cyan":    "6",
	"white":   "7",
	"default": "9",

	"txt": "3",
	"bg":  "4",

	"bold":      "1",
	"italic":    "3",
	"underline": "4",
	"strike":    "9",

	"escape":    "\033[",
	"separator": ";",
	"end":       "m",
}
