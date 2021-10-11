package scheduler

import (
	"context"
	"fmt"
	"sync"
)

// TaskScheduler is an entity for task scheduler
type TaskScheduler struct {
	mutex sync.Mutex
}

func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{}
}

func (ts *TaskScheduler) Run(ctx context.Context) {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	// call the service method

	fmt.Println("Hello World")
}
