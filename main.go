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
	fmt.Println("Start testing")

	tests := []engine.TestData{
		{
			Title:     "Small pool scheduling",
			InputFile: "input-data/smallpool.json",
		},
		{
			Title:     "Same time pool scheduling",
			InputFile: "input-data/sametimepool.json",
		},
		{
			Title:     "Big pool scheduling",
			InputFile: "input-data/bigpool.json",
		},
	}

	//startComputing := time.Now()
	//naiveLoad := engine.ComputeNaiveLoad(tasks, len(load))
	//duration := time.Since(startComputing)

	for i := range tests {
		f, err := os.Open(tests[i].InputFile)
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
		tests[i].Duration = time.Since(startComputing)
		tests[i].Load = s.Load()
	}

	if err := engine.PlotDoubleLoad(tests, "loads.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Complete. loads.png created.")
}
