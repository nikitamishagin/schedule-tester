package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"schedule-tester/internal/engine"
	"schedule-tester/pkg/scheduler"
	"schedule-tester/pkg/scheduler/v1"
)

func main() {
	f, err := os.Open("input.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var (
		tasks        []scheduler.Task
		runningTasks []scheduler.Task
	)
	if err := json.NewDecoder(f).Decode(&tasks); err != nil {
		log.Fatal(err)
	}

	s := v1.NewScheduler()
	for i := range tasks {
		s.AddTask(tasks[i], &runningTasks)
	}

	load := s.Load()

	naiveLoad := engine.ComputeNaiveLoad(tasks, len(load))

	if err := engine.PlotDoubleLoad(naiveLoad, load, "loads.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created: load.csv, loads.png")
}
