package v2

import (
	"math"

	"github.com/nikitamishagin/schedule-tester/pkg/scheduler"
)

type Scheduler struct {
	load []int // Array representing the load (number of tasks) at each time tick
}

// NewScheduler creates a scheduler instance with an empty load
func NewScheduler() *Scheduler {
	return &Scheduler{
		load: []int{},
	}
}

// Load returns the current load schedule (tasks per time tick)
func (s *Scheduler) Load() []int {
	return s.load
}

// Calculates the least common multiple (LCM) of two integers
func lcm(a, b int) int {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	return a * b / gcd(a, b)
}

// Calculates the least common multiple (LCM) of a slice of integers
func totalLCM(periods []int) int {
	result := periods[0]
	for _, p := range periods[1:] {
		result = lcm(result, p)
	}
	return result
}

// AddTask schedules a new task and returns the best starting time
func (s *Scheduler) AddTask(newTask scheduler.Task, runningTasks *[]scheduler.Task) int {
	maxTime := 0
	if len(*runningTasks) != 0 {
		// Calculate LCM of periods of all running tasks to set a scheduling horizon
		periods := make([]int, len(*runningTasks))
		for i, task := range *runningTasks {
			periods[i] = task.Period
		}
		maxTime = 1 * totalLCM(periods)
	} else {
		// If there are no running tasks, use the new task's period as scheduling horizon
		maxTime = newTask.Period + 1
	}

	bestScore := math.MaxInt     // Minimal load found (the lower, the better)
	bestStart := newTask.Arrived // Best start time for the new task found so far
	prevScore := 0               // Holds the score from the previous iteration to detect early stop condition

	// Try all possible start times within one period, beginning at task arrival time
	for start := newTask.Arrived; start <= newTask.Arrived+newTask.Period; start++ {
		score := 0 // Max load encountered for this start time
		for k := 0; ; k++ {
			tick := start + k*newTask.Period // Calculate next occurrence time for the task
			if tick >= maxTime {
				break // Stop if exceeded scheduling horizon
			}

			// Extend load array if current tick is out of bounds
			if len(s.load) < maxTime {
				s.load = append(s.load, make([]int, maxTime-len(s.load))...)
			}

			// Update score if load at this tick is higher than previous
			if s.load[tick] > score {
				score = s.load[tick]
			}
		}

		// If the current score is the best so far, record this start time
		if score < bestScore {
			bestScore = score
			bestStart = start
		}

		// Now update prevScore after using its value in comparison
		if score == prevScore {
			break // Stop checking further options if the score hasn't changed
		}
		prevScore = score
	}

	// Schedule the task at the best found start time and update the load array
	for k := 0; ; k++ {
		tick := bestStart + k*newTask.Period
		if tick >= maxTime {
			break
		}
		s.load[tick]++ // Increment load to reflect the new task at this time
	}

	// Add the new task to the list of running tasks
	*runningTasks = append(*runningTasks, newTask)
	return bestStart // Return the selected start time for the new task
}
