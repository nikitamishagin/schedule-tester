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

	var tasks []scheduler.Task
	if err := json.NewDecoder(f).Decode(&tasks); err != nil {
		log.Fatal(err)
	}

	periods := make([]int, len(tasks))
	for i, t := range tasks {
		periods[i] = t.Period
	}

	maxTime := 2 * v1.TotalLCM(periods)

	s := v1.NewScheduler(maxTime)
	for _, task := range tasks {
		s.AddTask(task)
	}

	load := s.Load()
	naiveLoad := engine.ComputeNaiveLoad(tasks, maxTime)

	csvFile, err := os.Create("load.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	fmt.Fprintln(csvFile, "time,load")
	for t := 0; t < maxTime; t++ {
		fmt.Fprintf(csvFile, "%d,%d\n", t, load[t])
	}

	if err := engine.PlotDoubleLoad(naiveLoad, load[:maxTime], "loads.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created: load.csv, loads.png")
}
