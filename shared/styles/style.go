package styles

import "fmt"

type OutputStyle struct {
	Text       string
	Background string
	Format     []string
	Reset      string
	ANSI       string
}

func (t *OutputStyle) New(txt, background string, styles []string) {

	t.Text += txt
	t.Background = background
	t.Format = styles

	t.ANSI, t.Reset = generateANSISequence(txt, background, styles)
}

func (t *OutputStyle) Style(s string) string {
	return t.ANSI + s + t.Reset
}

func (t *OutputStyle) ShowWithStyle(s string) {
	fmt.Print(t.ANSI + s)
	fmt.Println(t.Reset)
}
