package task

import (
	"testing"
)

func TestIsThisAStatus(t *testing.T) {
	arg := "blah"
	got := IsThisAStatus(arg)

	if got {
		t.Errorf("With argument '%s':\nExpected 'false'; Got 'true'", arg)
	}

	arg = "done"
	got = IsThisAStatus(arg)

	if !got {
		t.Errorf("With argument '%s':\nExpected 'true'; Got 'false'", arg)
	}
}

func TestIsThisAPriority(t *testing.T) {
	arg := "blah"
	got := IsThisAPriority(arg)

	if got {
		t.Errorf("With argument '%s':\nExpected 'false'; Got 'true'", arg)
	}

	arg = "low"
	got = IsThisAPriority(arg)

	if !got {
		t.Errorf("With argument '%s':\nExpected 'true'; Got 'false'", arg)
	}
}
