package storage

import (
	"fmt"
	"os"

	"github.com/leo-alvarenga/to-go/shared/task"
	"gopkg.in/yaml.v3"
)

func ReadFromYamlFile(filename string, taskSlice *[]task.Task, done chan bool) {
	content, err := os.ReadFile(filename)

	if err == nil {
		task := new([]task.Task)
		err = yaml.Unmarshal(content, task)

		if err == nil {
			for index, todo := range *task {
				todo.Id = fmt.Sprintf("%s-%d", filename, index)
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
