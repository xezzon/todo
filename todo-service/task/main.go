package task

import (
	"sync"

	"github.com/google/uuid"

	taskPb "todo-service/gen/todo/task"
)

type TaskStore struct {
	tasks map[string]*taskPb.Task
	mu    sync.RWMutex
}

func NewTaskStore() TaskStore {
	return TaskStore{
		tasks: make(map[string]*taskPb.Task),
	}
}

func (ts *TaskStore) Add(req *taskPb.AddTaskReq) *taskPb.Task {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	Id, _ := uuid.NewV7()
	task := &taskPb.Task{
		Id:      Id.String(),
		Content: req.Content,
	}
	ts.tasks[Id.String()] = task
	return task
}

func (ts *TaskStore) GetAll() []*taskPb.Task {
	ts.mu.RLock()
	defer ts.mu.RUnlock()

	tasks := make([]*taskPb.Task, 0, len(ts.tasks))
	for _, task := range ts.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (ts *TaskStore) Delete(id string) {
	ts.mu.Lock()
	defer ts.mu.Unlock()

	delete(ts.tasks, id)
}
