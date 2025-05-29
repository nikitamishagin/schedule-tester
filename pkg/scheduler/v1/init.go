package v1

import (
	"math"

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
	maxTime := 0
	if len(*runningTasks) != 0 {
		periods := make([]int, len(*runningTasks))
		for i, task := range *runningTasks {
			periods[i] = task.Period
		}

		maxTime = 2 * totalLCM(periods)
	} else {
		maxTime = newTask.Period + 1
	}

	bestScore := math.MaxInt
	bestStart := newTask.Arrived

	for start := newTask.Arrived; start <= newTask.Arrived+newTask.Period; start++ {
		score := 0
		for k := 0; ; k++ {
			tick := start + k*newTask.Period
			if tick >= maxTime {
				break
			}

			if len(s.load) < maxTime {
				s.load = append(s.load, make([]int, maxTime-len(s.load))...)
			}

			if s.load[tick] > score {
				score = s.load[tick]
			}
		}
		if score < bestScore {
			bestScore = score
			bestStart = start
		}
	}

	for k := 0; ; k++ {
		tick := bestStart + k*newTask.Period
		if tick >= maxTime {
			break
		}
		s.load[tick]++
	}

	*runningTasks = append(*runningTasks, newTask)
	return bestStart
}
