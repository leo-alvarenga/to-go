package ngops

import (
	"testing"
	"time"

	"github.com/leo-alvarenga/to-go/shared/task"
)

func getNewTaskSliceWithOneTask(title string) (*[]task.Task, error) {
	tasks := new([]task.Task)
	ts := task.Task{
		Title: "Test",
	}

	err := addTask(tasks, ts, false)

	return tasks, err
}

func TestAddTask(t *testing.T) {

	tasks, got := getNewTaskSliceWithOneTask("Test")
	after := len(*tasks)

	if got != nil {
		t.Errorf("Unable to add task. Error:\n%s\n", got.Error())
	}

	if after <= 0 {
		t.Errorf("Add operation error. Expect lenght of 1 got %d\n", after)
	} else {
		if (*tasks)[0].Title != "Test" {
			t.Errorf("Add operation worked with unexpected result\n."+
				"Expect to add task with title '%s'; Added task has '%s' title instead\n",
				"Title",
				(*tasks)[0].Title,
			)
		}
	}

	tasks = nil
}

func TestRemoveTask(t *testing.T) {
	tasks, err := getNewTaskSliceWithOneTask("Test")

	if err != nil {
		t.Errorf("Something went wrong while setting up the test. Error:\n%s\n", err.Error())
	}

	ts := task.Task{
		Title: "Test",
	}

	tasks, got := removeTask(tasks, ts)

	if got != nil {
		t.Errorf("Unable to remove task. Error:\n%s\n", got.Error())
	}

	after := len(*tasks)

	if after > 0 {
		t.Errorf("Remove operation error. Expect lenght of 0 got %d\n", after)
	}

	tasks = nil
}

func TestGetDateInToGosFmt(t *testing.T) {
	_, m, _ := time.Now().Date()

	got := getDateInToGosFmt(2022, m, 1)
	expect := m.String() + " 1st, 2022"

	if got != expect {
		t.Errorf("Expected '%s'; Got '%s'\n", expect, got)
	}
}
