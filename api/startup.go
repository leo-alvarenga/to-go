package api

import (
	"github.com/leo-alvarenga/to-go/shared"
	"github.com/leo-alvarenga/to-go/shared/config"
)

/* Executes all necessary steps to spin up the engine */
func StartupEngine(cfg *config.ConfigValue) {
	cfg.New()
	cfg.LoadFromYaml(config.ConfigFile)

	lowPriorityTasks = new([]shared.Task)
	mediumPriorityTasks = new([]shared.Task)
	highPriorityTasks = new([]shared.Task)

	retrieveTasks()
}
