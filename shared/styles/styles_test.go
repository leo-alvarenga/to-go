package styles

import "testing"

func TestNew(t *testing.T) {
	a, b := new(OutputStyle), new(OutputStyle)

	a.New("black", "", []string{"italic"})
	b.New("black", "", []string{"italic"})

	if b.ANSI != a.ANSI {
		t.Error("*OutputStyle.New() does not present the same behaviour when using the exact same arguments.")
	}

	b.New("white", "black", []string{})

	if b.ANSI == a.ANSI {
		t.Error("*ConfigValue.New() presents the same behavior, even with different arguments.")
	}
}

func TestStyle(t *testing.T) {
	a := new(OutputStyle)

	a.New("black", "", []string{"italic"})

	got := a.Style("test")
	expected := a.ANSI + "test" + a.Reset

	if got != expected {
		t.Errorf("Expected %s; Got %s\n", expected, got)
	}
}

func TestGenerateANSISequence(t *testing.T) {
	a, _ := generateANSISequence("black", "blue", []string{"bold"})
	b, _ := generateANSISequence("white", "yellow", []string{"bold", "italic", "underline", "strike"})

	if a != "\033[30;44;1m" {
		t.Errorf(
			"Text 'black'; Background 'blue'; Style 'bold'"+
				"Expected '[30;44;1m'; Got %s\n",
			a[1:],
		)
	} else if b != "\033[37;43;1;3;4;9m" {
		t.Errorf(
			"Text 'white'; Background 'yellow'; Style 'bold', 'italic', 'underline', 'strike'"+
				"Expected '[37;43;1;3;4;9m'; Got %s\n",
			b[1:],
		)
	}
}
