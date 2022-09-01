package engine

import (
	"fmt"
	"os"
	"sync"

	"github.com/leo-alvarenga/to-go/shared/task"
	"gopkg.in/yaml.v3"
)

func readFromFile(filename string, taskSlice *[]task.Task, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

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
			fmt.Println("couldn unmarshal", filename)
		}

		task = nil
	} else {
		file, _ := os.OpenFile(filename, os.O_CREATE, 0644)

		file.Close()
	}

	done <- true
}

/* Reads and retrieves tasks from each one of the task YAML files. */
func retrieveTasks() {
	var flags []chan bool
	var ref *[]task.Task
	wg := new(sync.WaitGroup)

	wg.Add(fileCount)

	for i, file := range taskFiles {
		flags = append(flags, make(chan bool))

		switch file[:2] {
		case "low":
			ref = lowPriorityTasks
		case "med":
			ref = mediumPriorityTasks
		default:
			ref = highPriorityTasks
		}

		go readFromFile(file, ref, flags[i], wg)
	}

	for _, f := range flags {
		<-f
	}

	wg.Wait()
}

/*
Returns pointers for each one of the dynamically allocated Task slices:
  - Pointer [0] -> Low priority tasks;
  - Pointer [1] -> Medium priority tasks;
  - Pointer [2] -> High priority tasks;
*/
func GetTasks() [fileCount]*[]task.Task {
	return [fileCount]*[]task.Task{
		highPriorityTasks, lowPriorityTasks, mediumPriorityTasks,
	}
}
