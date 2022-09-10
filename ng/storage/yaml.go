package storage

import (
	"errors"
	"os"

	"github.com/leo-alvarenga/to-go/shared/task"
	"gopkg.in/yaml.v3"
)

func readFromYamlFile(filename string, taskSlice *[]task.Task, done chan bool) {
	content, err := os.ReadFile(filename)

	if err == nil {
		tasks := new([]task.Task)
		err = yaml.Unmarshal(content, tasks)

		if err == nil {
			for _, todo := range *tasks {
				*taskSlice = append(*taskSlice, todo)
			}
		} else {
			done <- false
			return
		}

		tasks = nil
	} else {
		file, _ := os.OpenFile(filename, os.O_CREATE, 0644)

		file.Close()
	}

	done <- true
}

func WriteToYamlFile(filename string, taskSlice *[]task.Task, done chan error) {
	content, err := yaml.Marshal(taskSlice)

	if err == nil {
		os.WriteFile(filename, content, 0666)
	}

	done <- err
}

/* Reads and retrieves tasks from each one of the task YAML files. */
func RetrieveTasksFromYaml(taskFilenames [3]string, list *task.TaskList) error {

	low, med, high := make(chan bool, 1), make(chan bool, 1), make(chan bool, 1)

	for _, file := range taskFilenames {
		switch file[:3] {
		case "low":
			go readFromYamlFile(file, list.Low, low)
		case "med":
			go readFromYamlFile(file, list.Medium, med)
		default:
			go readFromYamlFile(file, list.High, high)
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

	list.SyncNextId()

	return nil
}
