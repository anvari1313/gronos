package scheduler

import (
	"os"
	"time"
)

var tm = &taskManager{
	tasks:           make(map[string]*Task),
	disabledTasks:   make(map[string]bool),
	taskLastRunTime: make(map[string]int64),
}

// StartLooper starts the main looper for tasks
func StartLooper() {

	path := os.Args[1]

	tasks, err := LoadTasksFromDir(tm, path)
	if err != nil {
		panic(err)
	}

	tm.initialize(tasks)

	ticker := time.Tick(1 * time.Second)

	go looper(tm.tasks, ticker)
}

func looper(tasks Tasks, ticker <-chan time.Time) {
	for {
		t := <-ticker
		for _, task := range tasks {
			if task.IsTime(&t) {
				task.Execute()
			}
		}
	}
}
