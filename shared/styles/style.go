package styles

type TextStyle struct {
	Text       string
	Background string
	Format     []string
	Reset      string
	ANSI       string
}

func (t *TextStyle) New(txt, background string, styles []string) {

	t.Text += txt
	t.Background = background
	t.Format = styles

	t.ANSI, t.Reset = generateANSISequence(txt, background, styles)
}
