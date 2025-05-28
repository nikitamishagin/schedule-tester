package engine

import (
	"schedule-tester/pkg/scheduler"
)

func ComputeNaiveLoad(tasks []scheduler.Task, maxTime int) []int {
	load := make([]int, maxTime)
	for _, t := range tasks {
		for tick := t.Arrived; tick < maxTime; tick += t.Period {
			load[tick]++
		}
	}
	return load
}
