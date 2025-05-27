package engine

import "math"

type Task struct {
	ID      int `json:"id"`
	Period  int `json:"period"`
	Arrived int `json:"arrived"`
}

type Scheduler interface {
	AddTask(t Task) int
	Load() []int
}

const (
	MaxTime  = 10000
	MaxLookK = 20
)

type GreedyScheduler struct {
	load []int
}

func NewGreedyScheduler() *GreedyScheduler {
	return &GreedyScheduler{
		load: make([]int, MaxTime),
	}
}

func (s *GreedyScheduler) AddTask(t Task) int {
	bestScore := math.MaxInt
	bestStart := t.Arrived

	for start := t.Arrived; start <= t.Arrived+t.Period; start++ {
		score := 0
		for k := 0; k <= MaxLookK; k++ {
			tick := start + k*t.Period
			if tick >= MaxTime {
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
		if tick >= MaxTime {
			break
		}
		s.load[tick]++
	}

	return bestStart
}

func (s *GreedyScheduler) Load() []int {
	return s.load
}
