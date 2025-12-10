package core

import (
	"sync"
	"time"
)

type TaskTracker struct {
	Tasks  []*Task `json:"tasks"`
	NextId int
}

var (
	instance *TaskTracker
	once     sync.Once
)

func GetTaskTracker() *TaskTracker {
	once.Do(func() {
		instance = &TaskTracker{
			Tasks:  make([]*Task, 0),
			NextId: 0,
		}
	})

	return instance
}

// add
func (t *TaskTracker) Add(desc string) {
	t.Tasks = append(t.Tasks, NewTask(t.NextId, desc))
}

// edit
func (t *TaskTracker) Update(id int, desc string) {
	if desc != "" {
		currentTask := t.Tasks[id]
		currentTask.Description = desc
		currentTask.CreatedAt = time.Now()
		currentTask.Status = PENDING
	}
}

// delete
func (t *TaskTracker) Delete(id int) {
	t.Tasks = append(t.Tasks[:id], t.Tasks[id+1:]...)
}

// mark as done
func (t *TaskTracker) MarkDone(id int) {
	currentTask := t.Tasks[id]

	currentTask.Status = COMPLETED
	currentTask.CompletedAt = time.Now()
}

// mark as progress
func (t *TaskTracker) MarkProgress(id int) {
	t.Tasks[id].Status = PROGRESS
}

// mark as todo
func (t *TaskTracker) MarkTodo(id int) {
	t.Tasks[id].Status = PENDING
}

// Get all tasks
func (t *TaskTracker) GetAll() []*Task {
	return t.Tasks
}

// Get all pending tasks
func (t *TaskTracker) GetPendingTask() []*Task {
	filter := make([]*Task, 0)
	for _, task := range t.Tasks {
		if task.Status == PENDING {
			filter = append(filter, task)
		}
	}

	return filter
}

// Get all progress task
func (t *TaskTracker) GetProgressTask() []*Task {
	filter := make([]*Task, 0)
	for _, task := range t.Tasks {
		if task.Status == PROGRESS {
			filter = append(filter, task)
		}
	}

	return filter
}

// Get all completed tasks
func (t *TaskTracker) GetCompletedTask() []*Task {
	filter := make([]*Task, 0)
	for _, task := range t.Tasks {
		if task.Status == COMPLETED {
			filter = append(filter, task)
		}
	}

	return filter
}

// update NextID
func (t *TaskTracker) UpdateNextID() int {
	maxId := 0
	for _, task := range t.Tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}

	return maxId + 1
}
