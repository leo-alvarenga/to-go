package task

import (
	"github.com/leo-alvarenga/to-go/shared/cfg"
	"github.com/leo-alvarenga/to-go/shared/styles"
)

type Task struct {
	Id          string
	Title       string
	Description string
	Priority    string
	Status      string
	CreatedIn   string
	FinishedIn  string
}

/*
 - title (required)
 - createdIn (autoset)
 - description (default="...")
 - priority (default=low)
 - dueTo (default=unset)
*/

func (t Task) GetStatusCharacterStyled(c cfg.StatusColor, useUnicode bool) string {
	style := new(styles.OutputStyle)
	chr := []string{}

	switch t.Status {
	case done:
		style.New(c.Done, "", []string{"bold"})
		chr = append(chr, "☑")
		chr = append(chr, "v")
	case pending:
		style.New(c.Pending, "", []string{"bold"})
		chr = append(chr, "?")
		chr = append(chr, " ")
	default:
		style.New(c.Doing, "", []string{"bold"})
		chr = append(chr, "☐")
		chr = append(chr, ".")
	}

	if useUnicode {
		return style.Style(chr[0])
	}

	return style.Style(chr[1])
}

func (t Task) GetPriorityCharacterStyled(c cfg.PriorityColor) string {
	style := new(styles.OutputStyle)
	chr := ""

	switch t.Priority {
	case high:
		style.New(c.High, "", []string{"bold"})
		chr = "^"
	case medium:
		style.New(c.Medium, "", []string{"bold"})
		chr = "-"
	default:
		style.New(c.Low, "", []string{"bold"})
		chr = "v"
	}

	return style.ANSI + chr + style.Reset
}
