package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"schedule-tester/internal/engine"
)

func lcm(a, b int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	return a * b / gcd(a, b)
}

func totalLCM(periods []int) int {
	result := periods[0]
	for _, p := range periods[1:] {
		result = lcm(result, p)
	}
	return result
}

func main() {
	f, err := os.Open("input.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var tasks []engine.Task
	if err := json.NewDecoder(f).Decode(&tasks); err != nil {
		log.Fatal(err)
	}

	s := engine.NewGreedyScheduler()

	for _, task := range tasks {
		s.AddTask(task)
	}

	load := s.Load()

	periods := make([]int, len(tasks))
	for i, t := range tasks {
		periods[i] = t.Period
	}
	cutoff := 2 * totalLCM(periods)
	if cutoff > len(load) {
		cutoff = len(load)
	}

	csvFile, err := os.Create("load.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	fmt.Fprintln(csvFile, "time,load")
	for t := 0; t < cutoff; t++ {
		fmt.Fprintf(csvFile, "%d,%d\n", t, load[t])
	}

	if err := engine.PlotLoad(load[:cutoff], "load.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Created: load.csv, load.png")
}
