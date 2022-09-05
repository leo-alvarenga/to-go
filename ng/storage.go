package ng

import (
	"errors"

	"github.com/leo-alvarenga/to-go/ng/storage"
	"github.com/leo-alvarenga/to-go/shared/task"
)

/* Reads and retrieves tasks from each one of the task YAML files. */
func retrieveTasksFromYaml() error {

	low, med, high := make(chan bool, 1), make(chan bool, 1), make(chan bool, 1)

	for _, file := range taskFilenames {
		switch file[:3] {
		case "low":
			go storage.ReadFromYamlFile(file, lowPriorityTasks, low)
		case "med":
			go storage.ReadFromYamlFile(file, mediumPriorityTasks, med)
		default:
			go storage.ReadFromYamlFile(file, highPriorityTasks, high)
		}
	}

	ok := (<-high && <-med && <-low)

	if !ok {
		return errors.New(
			"One or more files could not be unmarshaled," +
				"meaning they were found and read" +
				"but couldn't be interpreted.",
		)
	}

	return nil
}

/*
Returns pointers for each one of the dynamically allocated Task slices
*/
func GetTasks() map[string]*[]task.Task {
	return map[string]*[]task.Task{
		"high":   highPriorityTasks,
		"medium": mediumPriorityTasks,
		"low":    lowPriorityTasks,
	}
}
