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

func TestTaskListOps(t *testing.T) {
	task := Task{
		Title:       "ok",
		Description: "",
		Priority:    "",
	}

	list := new(TaskList)
	list.New()

	t.Log("Adding task with no priority...")
	err := list.Add(task)

	if err == nil {
		t.Error("Expected error when adding task with no priority; Got nil")
	}

	task.Priority = low
	t.Log("Adding task with priority...")
	err = list.Add(task)

	if err != nil {
		t.Errorf("Expected no error; Got '%s'\n", err.Error())
	}

	task.Status = done
	t.Log("Updating task status...")
	err = list.Update(task)

	if err != nil {
		t.Errorf("Expected no error; Got '%s'\n", err.Error())
	}

	if (*list.Low)[0].Status != done {
		t.Errorf("Expected status od 'done' after update; Got '%s'\n", (*list.Low)[0].Status)
	}

	t.Log("Removing task...")
	err = list.Remove(task.Title)

	if err != nil {
		t.Errorf("Expected no error; Got '%s'\n", err.Error())
	}

	if len(*list.Low) > 0 {
		t.Errorf("Expected length of 0 after removal; Got '%d'\n", len(*list.Low))
	}

	task.Priority = medium
	t.Log("Adding task with priority again...")
	err = list.Add(task)

	if err != nil {
		t.Errorf("Expected no error; Got '%s'\n", err.Error())
	}

	old := task
	task.Title = "a"
	task.Priority = high

	t.Log("Editing task priority and title...")
	err = list.Edit(old, task)

	if err != nil {
		t.Errorf("Expected no error; Got '%s'\n", err.Error())
	}

	if len(*list.Medium) > len(*list.High) {
		t.Errorf("Expected to change the task priority; Got '%s' instead", (*list.Medium)[0].Priority)
	} else if (*list.High)[0].Title != task.Title {
		t.Errorf("Expected to change the task title; Expected '%s'; Got '%s' instead", task.Title, (*list.High)[0].Title)
	}
}
