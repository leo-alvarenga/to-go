package ngops

import (
	"testing"
	"time"

	"github.com/leo-alvarenga/to-go/shared/task"
)

func TestAddTask(t *testing.T) {
	tasks := new([]task.Task)
	ts := task.Task{
		Title: "okok",
	}

	before := len(*tasks)
	got := addTask(tasks, ts, false)
	after := len(*tasks)

	if got != nil {
		t.Errorf("Unable to add task. Error:\n%s\n", got.Error())
	}

	if after <= before {
		t.Errorf("Add operation error. Expect lenght of %d got %d\n", before+1, after)
	}
}

func TestGetDateInToGosFmt(t *testing.T) {
	_, m, _ := time.Now().Date()

	got := getDateInToGosFmt(2022, m, 1)
	expect := m.String() + " 1st, 2022"

	if got != expect {
		t.Errorf("Expected '%s'; Got '%s'\n", expect, got)
	}
}
