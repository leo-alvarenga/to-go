package task

import "github.com/leo-alvarenga/to-go/shared/styles"

type Task struct {
	Id          string
	Title       string
	Description string
	Priority    string
	Status      string
}

/*
 - title (required)
 - createdIn (autoset)
 - description (default="...")
 - priority (default=low)
 - dueTo (default=unset)
*/

func (t Task) GetStatusCharacterStyled() string {
	style := new(styles.OutputStyle)
	chr := ""

	switch t.Status {
	case done:
		style.New("green", "", []string{"bold"})
		chr = "☑"
	case pending:
		style.New("yellow", "", []string{"bold"})
		chr = "?"
	default:
		style.New("blue", "", []string{"bold"})
		chr = "☐"
	}

	return style.ANSI + chr + style.Reset
}

func (t Task) GetPriorityCharacterStyled() string {
	style := new(styles.OutputStyle)
	chr := ""

	switch t.Priority {
	case high:
		style.New("red", "", []string{"bold"})
		chr = "^"
	case medium:
		style.New("yellow", "", []string{"bold"})
		chr = "-"
	default:
		style.New("green", "", []string{"bold"})
		chr = "v"
	}

	return style.ANSI + chr + style.Reset
}
