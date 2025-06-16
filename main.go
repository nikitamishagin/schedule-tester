package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nikitamishagin/schedule-tester/internal/engine"
	"github.com/nikitamishagin/schedule-tester/pkg/scheduler"
	v1 "github.com/nikitamishagin/schedule-tester/pkg/scheduler/v1"
	v2 "github.com/nikitamishagin/schedule-tester/pkg/scheduler/v2"
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

	// Define available scheduler versions and their human-friendly names
	versions := []struct {
		Name          string
		MakeScheduler func() scheduler.Scheduler // suppose all Schedulers implement this
	}{
		{"v1", func() scheduler.Scheduler { return v1.NewScheduler() }},
		{"v2", func() scheduler.Scheduler { return v2.NewScheduler() }},
	}

	// Outer loop: over all input-data test cases
	for i := range tests {
		f, err := os.Open(tests[i].InputFile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		var tasks []scheduler.Task
		if err := json.NewDecoder(f).Decode(&tasks); err != nil {
			log.Fatal(err)
		}

		maxTime := 1200
		tests[i].NaiveLoad = engine.ComputeNaiveLoad(tasks, maxTime)

		// Store all results for all algorithms
		tests[i].Loads = make(map[string][]int)
		tests[i].Durations = make(map[string]time.Duration)

		for _, ver := range versions {
			s := ver.MakeScheduler()
			var runningTasks []scheduler.Task
			startComputing := time.Now()
			for j := range tasks {
				s.AddTask(tasks[j], &runningTasks)
			}
			duration := time.Since(startComputing)
			tests[i].Loads[ver.Name] = s.Load()
			tests[i].Durations[ver.Name] = duration
		}
	}

	// Pass all collected data for all algorithms in all tests to plotting function
	if err := engine.PlotMultiRowLoads(tests, "loads.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Complete. loads.png created.")
}
