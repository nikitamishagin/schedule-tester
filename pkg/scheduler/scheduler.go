package scheduler

type Task struct {
	ID      int `json:"id"`
	Period  int `json:"period"`
	Arrived int `json:"arrived"`
}

type Scheduler interface {
	AddTask(Task) int
	Load() []int
}
