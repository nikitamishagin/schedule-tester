package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

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

	startComputing := time.Now()
	for i := range tasks {
		s.AddTask(tasks[i], &runningTasks)
	}
	duration := time.Since(startComputing)
	load := s.Load()

	scheduledLoadPlot := engine.LoadPlot{
		Load:     load,
		Duration: duration,
	}

	startComputing = time.Now()
	naiveLoad := engine.ComputeNaiveLoad(tasks, len(load))
	duration = time.Since(startComputing)

	naiveLoadPlot := engine.LoadPlot{
		Load:     naiveLoad,
		Duration: duration,
	}
	if err := engine.PlotDoubleLoad(naiveLoadPlot, scheduledLoadPlot, "loads.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created: load.csv, loads.png")
}
