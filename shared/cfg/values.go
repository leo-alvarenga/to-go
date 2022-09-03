package cfg

/* The max length a line can have when being printed to stdout */
const DefaultMaxLineLen int = 59

const colorReset string = "\033[0m"

var DefaultColors = ColorScheme{
	Title:   "cyan",
	Accent:  "purple",
	Error:   "red",
	Warning: "yellow",
	Success: "green",
	Reset:   colorReset,
}
