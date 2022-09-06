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
		task := new([]task.Task)
		err = yaml.Unmarshal(content, task)

		if err == nil {
			for _, todo := range *task {
				*taskSlice = append(*taskSlice, todo)
			}
		} else {
			done <- false
			return
		}

		task = nil
	} else {
		file, _ := os.OpenFile(filename, os.O_CREATE, 0644)

		file.Close()
	}

	done <- true
}

func WriteToYamlFile(filename string, taskSlice *[]task.Task) error {
	content, err := yaml.Marshal(taskSlice)

	if err == nil {
		os.WriteFile(filename, content, 0666)
	}

	return err
}

/* Reads and retrieves tasks from each one of the task YAML files. */
func RetrieveTasksFromYaml(taskFilenames [3]string, h, m, l *[]task.Task) error {

	low, med, high := make(chan bool, 1), make(chan bool, 1), make(chan bool, 1)

	for _, file := range taskFilenames {
		switch file[:3] {
		case "low":
			go readFromYamlFile(file, l, low)
		case "med":
			go readFromYamlFile(file, m, med)
		default:
			go readFromYamlFile(file, h, high)
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
