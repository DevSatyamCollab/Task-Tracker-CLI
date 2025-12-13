package service

import (
	"todo-list/internal/core"
	"todo-list/internal/storage"
)

type TaskService struct {
	tracker *core.TaskTracker
	storage *storage.Storage
}

func NewTaskService(t *core.TaskTracker, s *storage.Storage) *TaskService {
	return &TaskService{tracker: t, storage: s}
}

// add Task
func (ts *TaskService) AddTask(desc string) error {
	ts.tracker.Add(desc)

	if err := ts.storage.Save(ts.tracker); err != nil {
		return err
	}

	return nil
}

// update task
func (ts *TaskService) UpdateTask(id int, desc string) error {
	index, err := ts.FindTask(id)
	if err != nil {
		return err
	}

	ts.tracker.Update(index, desc)

	if err := ts.storage.Save(ts.tracker); err != nil {
		return err
	}

	return nil
}

// delete task
func (ts *TaskService) DeleteTask(id int) error {
	index, err := ts.FindTask(id)
	if err != nil {
		return err
	}

	ts.tracker.Delete(index)

	if err := ts.storage.Save(ts.tracker); err != nil {
		return err
	}

	return nil
}

// mark as done
func (ts *TaskService) MarkDone(id int) error {
	index, err := ts.FindTask(id)
	if err != nil {
		return err
	}

	ts.tracker.MarkDone(index)

	if err := ts.storage.Save(ts.tracker); err != nil {
		return err
	}

	return nil
}

// mark as progress
func (ts *TaskService) MarkProgress(id int) error {
	index, err := ts.FindTask(id)
	if err != nil {
		return err
	}

	ts.tracker.MarkProgress(index)

	if err := ts.storage.Save(ts.tracker); err != nil {
		return err
	}

	return nil
}

// mark as todo
func (ts *TaskService) MarkTodo(id int) error {
	index, err := ts.FindTask(id)
	if err != nil {
		return err
	}

	ts.tracker.MarkTodo(index)

	if err := ts.storage.Save(ts.tracker); err != nil {
		return err
	}

	return nil
}

func (ts *TaskService) GetAllTask() []*core.Task {
	return ts.tracker.GetAll()
}

func (ts *TaskService) GetAllPendingTask() []*core.Task {
	return ts.tracker.GetPendingTask()
}

func (ts *TaskService) GetAllCompletedTask() []*core.Task {
	return ts.tracker.GetCompletedTask()
}

func (ts *TaskService) GetAllProgressTask() []*core.Task {
	return ts.tracker.GetProgressTask()
}

// find task
func (ts *TaskService) FindTask(id int) (int, error) {
	tasklist := ts.tracker.GetAll()

	for i, task := range tasklist {
		if id == task.Id {
			return i, nil
		}
	}

	return -1, core.ErrTaskNotFound
}
