package task

import (
	"fmt"

	"github.com/leo-alvarenga/to-go/shared/cfg"
	"github.com/leo-alvarenga/to-go/shared/styles"
)

type Task struct {
	Id          int
	Title       string
	Description string
	Priority    string
	Status      string
	CreatedIn   string
	FinishedIn  string
}

func (t Task) GetStatusCharacterStyled(c cfg.StatusColor, useUnicode, withLabel bool) string {
	style := new(styles.OutputStyle)
	chr := []string{}

	switch t.Status {
	case done:
		style.New(c.Done, "", []string{"bold"})

		if withLabel {
			chr = append(chr, done+" ☑")
			chr = append(chr, done+" v")
		} else {
			chr = append(chr, "☑")
			chr = append(chr, "v")
		}
	case pending:
		style.New(c.Pending, "", []string{"bold"})

		if withLabel {
			chr = append(chr, pending+" ?")
			chr = append(chr, pending+" ")
		} else {
			chr = append(chr, "?")
			chr = append(chr, " ")
		}
	default:
		style.New(c.Doing, "", []string{"bold"})

		if withLabel {
			chr = append(chr, doing+" ☐")
			chr = append(chr, doing+" .")
		} else {
			chr = append(chr, "☐")
			chr = append(chr, ".")
		}
	}

	if useUnicode {
		return style.Style(chr[0])
	}

	return style.Style(chr[1])
}

func (t Task) GetPriorityCharacterStyled(c cfg.PriorityColor, withLabel bool) string {
	style := new(styles.OutputStyle)
	chr := ""

	switch t.Priority {
	case high:
		style.New(c.High, "", []string{"bold"})
		if withLabel {
			chr = high + " ^"
		} else {
			chr = "^"
		}
	case medium:
		style.New(c.Medium, "", []string{"bold"})

		if withLabel {
			chr = medium + " -"
		} else {
			chr = "-"
		}
	default:
		style.New(c.Low, "", []string{"bold"})
		if withLabel {
			chr = low + " v"
		} else {
			chr = "v"
		}
	}

	return style.ANSI + chr + style.Reset
}

func (t Task) DisplayTask(c cfg.ConfigValue) {
	style := new(styles.OutputStyle)
	out := ""

	style.New(c.Colors.Attention, "", []string{"bold"})
	out += style.Style(" > ")

	out += "[" + t.GetPriorityCharacterStyled(c.Colors.Priority, false) + "] "
	out += "[" + t.GetStatusCharacterStyled(c.Colors.Status, c.UseUnicode, false) + "] "

	if len(t.Title) <= 21 {
		out += style.Style(t.Title)
	} else {
		out += style.Style(t.Title[:20])
	}

	out += "\t\t(" + t.CreatedIn + " - "

	if len(t.FinishedIn) < 1 {
		out += "[...]"
	} else {
		out += t.FinishedIn
	}

	out += ")"

	fmt.Println(out)
}

func (t Task) Describe(c cfg.ConfigValue) {
	style := new(styles.OutputStyle)
	style.New(c.Colors.Attention, "", []string{"bold", "underline"})

	style.ShowWithStyle("\n" + t.Title)
	style.New("", "", []string{"bold"})

	style.ShowWithStyle(t.GetPriorityCharacterStyled(c.Colors.Priority, true))
	style.ShowWithStyle(t.GetStatusCharacterStyled(c.Colors.Status, c.UseUnicode, true))

	date := "(" + t.CreatedIn + " - "

	if t.FinishedIn == "" {
		date += "[...]"
	} else {
		date += t.FinishedIn
	}

	date += ")"

	style.New(c.Colors.Attention, "", []string{"bold"})
	style.ShowWithStyle(date)

	style.New("", "", []string{})
	if len(t.Description) > 40 {
		var i int

		for i = 0; i+40 < len(t.Description); i += 40 {
			style.ShowWithStyle(t.Description[i : i+40])
		}

		style.ShowWithStyle(t.Description[i:])
	} else if len(t.Description) == 0 {
		style.ShowWithStyle("[Empty]")
	} else {
		style.ShowWithStyle("\n" + t.Description)
	}
}
