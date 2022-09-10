package ngops

import (
	"fmt"
	"time"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/ng/storage"
)

func getDateInToGosFmt(y int, m time.Month, d int) string {
	extra := ""

	switch d % 10 {
	case 1:
		extra += "st"
	case 2:
		extra += "nd"
	case 3:
		extra += "rd"
	default:
		extra += "th"
	}

	return fmt.Sprintf("%s %d%s, %d", m.String(), d, extra, y)
}

func writeToYamlWrapper() error {
	var err [3]chan error

	for i := 0; i < 3; i++ {
		err[i] = make(chan error)
	}

	go storage.WriteToYamlFile(ng.TaskFilenamesMapped["high"], ng.TaskList.High, err[0])
	go storage.WriteToYamlFile(ng.TaskFilenamesMapped["medium"], ng.TaskList.Medium, err[1])
	go storage.WriteToYamlFile(ng.TaskFilenamesMapped["low"], ng.TaskList.Low, err[2])

	for _, e := range err {
		status := <-e

		if status != nil {
			return status
		}
	}

	return nil
}
