package ngops

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/shared/cfg"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func Dashboard(config *cfg.ConfigValue) {
	http.HandleFunc("/", landingPage)

	// ng calls
	http.HandleFunc("/api/get", getTasksHandler)
	http.HandleFunc("/api/add", addTaskHandler)
	http.HandleFunc("/api/remove", removeTaskHandler)
	http.HandleFunc("/api/edit", editTaskHandler)
	http.HandleFunc("/api/update", updateTaskHandler)

	log.Println("Serving dashboard on port " + config.DashboardPort)
	http.ListenAndServe(":"+config.DashboardPort, nil)
}

func landingPage(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		ex, _ := os.Executable()
		http.ServeFile(w, req, filepath.Join(filepath.Dir(ex), "static/index.html"))
	}
}

func getTasksHandler(w http.ResponseWriter, req *http.Request) {
	out, _ := json.Marshal([][]task.Task{*ng.TaskList.High, *ng.TaskList.Medium, *ng.TaskList.Low})

	fmt.Fprint(w, out)
}

func addTaskHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "ok")
}

func removeTaskHandler(w http.ResponseWriter, req *http.Request) {

}

func editTaskHandler(w http.ResponseWriter, req *http.Request) {

}

func updateTaskHandler(w http.ResponseWriter, req *http.Request) {

}
