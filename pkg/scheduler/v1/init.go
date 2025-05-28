package v1

import (
	"math"

	"schedule-tester/pkg/scheduler"
)

type Scheduler struct {
	maxTime int
	load    []int
}

func NewScheduler(maxTime int) *Scheduler {
	return &Scheduler{
		maxTime: maxTime,
		load:    make([]int, maxTime),
	}
}

func (s *Scheduler) MaxTime() int {
	return s.maxTime
}

func (s *Scheduler) Load() []int {
	return s.load
}

func LCM(a, b int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	return a * b / gcd(a, b)
}

func TotalLCM(periods []int) int {
	result := periods[0]
	for _, p := range periods[1:] {
		result = LCM(result, p)
	}
	return result
}

func (s *Scheduler) AddTask(t scheduler.Task) int {
	bestScore := math.MaxInt
	bestStart := t.Arrived

	for start := t.Arrived; start <= t.Arrived+t.Period; start++ {
		score := 0
		for k := 0; ; k++ {
			tick := start + k*t.Period
			if tick >= s.maxTime {
				break
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
		tick := bestStart + k*t.Period
		if tick >= s.maxTime {
			break
		}
		s.load[tick]++
	}

	return bestStart
}
