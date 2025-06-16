package engine

import (
	"time"

	"github.com/nikitamishagin/schedule-tester/pkg/scheduler"
)

type TestData struct {
	Title     string
	InputFile string
	NaiveLoad []int
	Load      []int
	Duration  time.Duration
}

func ComputeNaiveLoad(tasks []scheduler.Task, maxTime int) []int {
	load := make([]int, maxTime)
	for _, t := range tasks {
		for tick := t.Arrived; tick < maxTime; tick += t.Period {
			load[tick]++
		}
	}
	return load
}
