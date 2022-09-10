package task

import (
	"errors"
)

/* A container that holds all the tasks in dynamicaly allocated slices */
type TaskList struct {
	Low    *[]Task
	Medium *[]Task
	High   *[]Task
	nextId int
}

/* Initializes a newly allocated TaskList */
func (t *TaskList) New() {
	t.Low = new([]Task)
	t.Medium = new([]Task)
	t.High = new([]Task)

	t.nextId = 0
}

/* Updates the value meant to be the next task id */
func (t *TaskList) SyncNextId() {
	t.nextId = (len(*t.High) + len(*t.Medium) + len(*t.Low)) - 3

	if t.nextId < 0 {
		t.nextId = 0
	}
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

/* Returns the task priotity if it exists and its index in the task slice, or an empty string and -1 if it doesn't */
func (t *TaskList) doesThisTaskExists(title string) (string, int) {
	for i, item := range *t.High {
		if item.Title == title {
			return high, i
		}
	}

	for i, item := range *t.Medium {
		if item.Title == title {
			return medium, i
		}
	}

	for i, item := range *t.Low {
		if item.Title == title {
			return low, i
		}
	}

	return "", -1
}

func (t *TaskList) GetTaskByTitle(title string) (Task, error) {
	priority, index := t.doesThisTaskExists(title)

	if !IsThisAPriority(priority) {
		return Task{}, errors.New("Task '" + title + "' not found.")
	}

	var task Task

	switch priority {
	case high:
		task = (*t.High)[index]
	case medium:
		task = (*t.Medium)[index]
	case low:
		task = (*t.Low)[index]
	}

	return task, nil
}

/* Finds the index of the 'ts' Task; Returns an index of -1 and an error if it could not be found  */
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

/* Adds a 'ts' Task to its corresponding Task slice */
func (t *TaskList) Add(ts Task) error {
	if !IsThisAPriority(ts.Priority) {
		return errors.New("Task does not have a valid priority.")
	}

	p, _ := t.doesThisTaskExists(ts.Title)
	if IsThisAPriority(p) {
		return errors.New("There is already a task with the title '" + ts.Title + "'.")
	}

	switch ts.Priority {
	case high:
		ts.Id = t.nextId
		(*t.High) = append((*t.High), ts)
	case medium:
		ts.Id = t.nextId
		(*t.Medium) = append((*t.Medium), ts)
	case low:
		ts.Id = t.nextId
		(*t.Low) = append((*t.Low), ts)
	}

	t.nextId++

	return nil
}

/* Removes a 'ts' Task from its corresponding Task slice */
func (t *TaskList) Remove(title string) error {
	priority, index := t.doesThisTaskExists(title)

	if !IsThisAPriority(priority) || index < 0 {
		return errors.New("There are no tasks with the title '" + title + "'.")
	}

	switch priority {
	case high:
		(*t.High)[index] = (*t.High)[len(*t.High)-1]
		(*t.High) = (*t.High)[:len(*t.High)-1]
	case medium:
		(*t.Medium)[index] = (*t.Medium)[len(*t.Medium)-1]
		(*t.Medium) = (*t.Medium)[:len(*t.Medium)-1]
	case low:
		(*t.Low)[index] = (*t.Low)[len(*t.Low)-1]
		(*t.Low) = (*t.Low)[:len(*t.Low)-1]
	}

	return nil
}

/* Edits a task */
func (t *TaskList) Edit(old, ts Task) error {
	priority, index := t.doesThisTaskExists(old.Title)

	if !IsThisAPriority(priority) || index < 0 {
		return errors.New("There are no tasks with the title '" + old.Title + "'.")
	}

	if old.Priority == ts.Priority {
		switch priority {
		case high:
			(*t.High)[index] = ts
		case medium:
			(*t.Medium)[index] = ts
		case low:
			(*t.Low)[index] = ts
		}
	} else {
		if old.Title != ts.Title {
			priority, index = t.doesThisTaskExists(ts.Title)
			if IsThisAPriority(priority) || index >= 0 {
				return errors.New("There is already a task with the title '" + ts.Title + "'.")
			}
		}

		err := t.Remove(old.Title)
		if err != nil {
			return err
		}

		err = t.Add(ts)
		if err != nil {
			return err
		}
	}

	return nil
}

/* Updates the status field of a Task */
func (t *TaskList) Update(ts Task) error {
	priority, index := t.doesThisTaskExists(ts.Title)

	if !IsThisAPriority(priority) || index < 0 {
		return errors.New("There are no tasks with the title '" + ts.Title + "'.")
	}

	switch priority {
	case high:
		(*t.High)[index].Status = ts.Status
		(*t.High)[index].FinishedIn = ts.FinishedIn
	case medium:
		(*t.Medium)[index].Status = ts.Status
		(*t.Medium)[index].FinishedIn = ts.FinishedIn
	case low:
		(*t.Low)[index].Status = ts.Status
		(*t.Low)[index].FinishedIn = ts.FinishedIn
	}

	return nil
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
