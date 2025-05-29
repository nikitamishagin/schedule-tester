package scheduler

type Task struct {
	ID      int `json:"id"`      // ID of the task.
	Period  int `json:"period"`  // Period is time for the task to be run again.
	Arrived int `json:"arrived"` // Arrived is time when the task arrived.
}

type Scheduler interface {
	AddTask(Task, []Task) int
	Load() []int
}
