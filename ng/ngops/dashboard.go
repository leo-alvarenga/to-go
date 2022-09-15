package ngops

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/leo-alvarenga/to-go/ng"
	"github.com/leo-alvarenga/to-go/shared/cfg"
	"github.com/leo-alvarenga/to-go/shared/task"
)

func openBrowser(port string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default:
		cmd = "xdg-open"
	}
	args = append(args, "localhost:"+port)
	return exec.Command(cmd, args...).Start()
}

func Dashboard(config *cfg.ConfigValue) {
	ex, _ := os.Executable()
	statics := filepath.Join(filepath.Dir(ex), "static/")

	// static files
	http.Handle("/", http.FileServer(http.Dir(statics)))

	// ng calls
	http.HandleFunc("/api/get", getTasksHandler)
	http.HandleFunc("/api/add", addTaskHandler)
	http.HandleFunc("/api/remove", removeTaskHandler)
	http.HandleFunc("/api/edit", editTaskHandler)
	http.HandleFunc("/api/update", updateTaskHandler)

	if config.OpenBrowser {
		go openBrowser(config.DashboardPort)
	}

	log.Println("Serving dashboard on port " + config.DashboardPort)
	http.ListenAndServe(":"+config.DashboardPort, nil)
}

func getTasksHandler(w http.ResponseWriter, req *http.Request) {
	res, err := json.Marshal(map[string][]task.Task{"content": ng.TaskList.GetAllTasks()})

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"content": []}`))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(res)
	}
}

func addTaskHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	defer req.Body.Close()

	t := new(task.Task)

	err := json.NewDecoder(req.Body).Decode(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	err = Add(*t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	killData()
	w.WriteHeader(http.StatusCreated)
}

func removeTaskHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodDelete {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	title := req.URL.Query().Get("title")

	if title == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"Invalid title"}`))
		return
	}

	err := Remove(title)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	killData()
	w.WriteHeader(http.StatusOK)
}

func editTaskHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPatch {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	defer req.Body.Close()

	t := new([]task.Task)

	err := json.NewDecoder(req.Body).Decode(t)
	if err != nil || len(*t) <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	err = Edit((*t)[1], (*t)[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	killData()
	w.WriteHeader(http.StatusOK)
}

func updateTaskHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPatch {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	defer req.Body.Close()

	t := new(task.Task)

	err := json.NewDecoder(req.Body).Decode(t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	err = Update(*t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"status":"` + err.Error() + `"}`))
		return
	}

	killData()
	w.WriteHeader(http.StatusOK)
}
