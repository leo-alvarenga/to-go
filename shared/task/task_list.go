package task

import (
	"errors"
)

/* A container that holds all the tasks in dynamicaly allocated slices */
type TaskList struct {
	Low    *[]Task
	Medium *[]Task
	High   *[]Task
}

/* Initializes a newly allocated TaskList */
func (t *TaskList) New() {
	t.Low = new([]Task)
	t.Medium = new([]Task)
	t.High = new([]Task)
}

/* Empty one of the slices, based on the 'priority' parameter */
func (t *TaskList) clearList(priority string) error {
	if !IsThisAPriority(priority) {
		return errors.New("Task does not have a valid priority.")
	}

	switch priority {
	case high:
		t.High = nil
		t.High = new([]Task)
	case medium:
		t.Medium = nil
		t.Medium = new([]Task)
	case low:
		t.Low = nil
		t.Low = new([]Task)
	}

	return nil
}

/* Checks whether or not a task 'ts' exists within one of the Task slices */
func (t *TaskList) doesThisTaskExists(ts Task) bool {
	for _, item := range *t.High {
		if item.Title == ts.Title {
			return true
		}
	}

	for _, item := range *t.Medium {
		if item.Title == ts.Title {
			return true
		}
	}

	for _, item := range *t.Low {
		if item.Title == ts.Title {
			return true
		}
	}

	return false
}

/* Finds the index of the 'ts' Task; Returns an index of -1 and an error if it could not be found  */
// todo -> make so ids aare unique, even in yaml mode
func (t *TaskList) FindIndex(ts Task) (int, error) {
	if !IsThisAPriority(ts.Priority) {
		return -1, errors.New("Task does not have a valid priority.")
	}

	var ref *[]Task
	switch ts.Priority {
	case high:
		ref = t.High
	case medium:
		ref = t.Medium
	case low:
		ref = t.Low
	}

	for i, item := range *ref {
		if item.Title == ts.Title {
			return i, nil
		}
	}

	return -1, errors.New("Task '" + ts.Title + "' not found.")
}

/* Finds the index of the Task using only its title as info; Returns an index of -1 and an error if it could not be found  */
func (t *TaskList) FindIndexByTitle(title string) (int, string, error) {
	for i, item := range *t.High {
		if item.Title == title {
			return i, item.Priority, nil
		}
	}

	for i, item := range *t.Medium {
		if item.Title == title {
			return i, item.Priority, nil
		}
	}

	for i, item := range *t.Low {
		if item.Title == title {
			return i, item.Priority, nil
		}
	}

	return -1, "", errors.New("Task '" + title + "' not found.")
}

/*
Gets the address of a Tasks by searching its index in the corresponding Task slice
Returns nil and an error if not found
*/
func (t *TaskList) GetTaskByIndex(index int, priority string) (*Task, error) {
	if !IsThisAPriority(priority) {
		return nil, errors.New("Task does not have a valid priority.")
	}

	var ref *[]Task
	switch priority {
	case high:
		ref = t.High
	case medium:
		ref = t.Medium
	case low:
		ref = t.Low
	}

	if index < len(*ref) && index >= 0 {
		return &(*ref)[index], nil
	}

	return nil, errors.New("Task not found.")
}

/*
Gets the address of a Tasks by searching its title in the corresponding Task slice
Returns nil and an error if not found
*/
func (t *TaskList) GetTaskByTitle(title string) (*Task, error) {
	for _, item := range *t.High {
		if item.Title == title {
			return &item, nil
		}
	}

	for _, item := range *t.Medium {
		if item.Title == title {
			return &item, nil
		}
	}

	for _, item := range *t.Low {
		if item.Title == title {
			return &item, nil
		}
	}

	return nil, errors.New("Task '" + title + "' not found.")
}

/* Adds a 'ts' Task to its corresponding Task slice */
func (t *TaskList) Add(ts Task) error {
	if !IsThisAPriority(ts.Priority) {
		return errors.New("Task does not have a valid priority.")
	}

	switch ts.Priority {
	case high:
		ts.Id = len((*t.High))
		(*t.High) = append((*t.High), ts)
	case medium:
		ts.Id = len((*t.Medium))
		(*t.Medium) = append((*t.Medium), ts)
	case low:
		ts.Id = len((*t.Low))
		(*t.Low) = append((*t.Low), ts)
	}

	return nil
}

/* Removes a 'ts' Task from its corresponding Task slice */
func (t *TaskList) Remove(ts Task) error {
	return nil
}

/* Edits a task */
func (t *TaskList) Edit(old, ts Task) error {
	return nil
}

/* Updates the status field of a Task */
func (t *TaskList) Update(ts Task) error {
	return t.Edit(ts, ts)
}

/* Gets an slice containing all the Task titles currently in use */
func (t *TaskList) GetAllTitles() (titles []string) {
	for _, task := range *t.High {
		titles = append(titles, task.Title)
	}

	for _, task := range *t.Medium {
		titles = append(titles, task.Title)
	}

	for _, task := range *t.Low {
		titles = append(titles, task.Title)
	}

	return
}

/* Gets an slice containing all the Tasks */
func (t *TaskList) GetAllTasks() (tasks []Task) {
	for _, task := range *t.High {
		tasks = append(tasks, task)
	}

	for _, task := range *t.Medium {
		tasks = append(tasks, task)
	}

	for _, task := range *t.Low {
		tasks = append(tasks, task)
	}

	return
}
