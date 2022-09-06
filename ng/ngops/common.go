package ngops

import (
	"fmt"
	"time"
)

func getDateInToGosFmt(y int, m time.Month, d int) string {
	extra := ""

	switch d % 10 {
	case 1:
		extra += "st"
	case 2:
		extra += "nd"
	case 3:
		extra += "rd"
	default:
		extra += "th"
	}

	return fmt.Sprintf("%s %d%s, %d", m.String(), d, extra, y)
}
