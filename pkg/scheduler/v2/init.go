package v2

import (
	"schedule-tester/pkg/scheduler"
)

type Scheduler struct {
	load []int
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		load: []int{},
	}
}

func (s *Scheduler) Load() []int {
	return s.load
}

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

func (s *Scheduler) AddTask(newTask scheduler.Task, runningTasks *[]scheduler.Task) int {
	// TODO: Improve logic with binary tree

	return 0
}
