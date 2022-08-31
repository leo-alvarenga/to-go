package api

import (
	"fmt"
	"os"

	"github.com/leo-alvarenga/to-go/shared"
	"gopkg.in/yaml.v3"
)

/* Reads and retrieves tasks from each one of the task YAML files. */
func retrieveTasks() {
	var task *[]shared.Task
	var ref *[]shared.Task

	for i, file := range taskFiles {
		f, err := os.ReadFile(file)

		if err != nil {
			panic(fmt.Sprintf("Missing \"%s\". Aborting...", file))
		}

		task = new([]shared.Task)
		err = yaml.Unmarshal(f, task)

		if err != nil {
			panic(fmt.Sprintf("Could not unpack \"%s\". Aborting...", file))
		}

		switch i {
		case 1:
			ref = &*mediumPriorityTasks
		case 2:
			ref = &*highPriorityTasks
		default:
			ref = &*lowPriorityTasks
		}

		for index, todo := range *task {
			todo.Id = fmt.Sprintf("%s-%d", file, index)
			*ref = append(*ref, todo)
		}

		task = nil
		ref = nil
	}
}

/*
Returns pointers for each one of the dynamically allocated Task slices:
  - Pointer [0] -> Low priority tasks;
  - Pointer [1] -> Medium priority tasks;
  - Pointer [2] -> High priority tasks;
*/
func GetTasks() [fileCount]*[]shared.Task {
	return [fileCount]*[]shared.Task{
		lowPriorityTasks, mediumPriorityTasks, highPriorityTasks,
	}
}
