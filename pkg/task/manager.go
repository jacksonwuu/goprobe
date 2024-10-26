package task

import "sync"

// TaskManager manage a bunch of probe task,
// it will start do probe for these probe jobs.
type TaskManager struct {
	tasks map[string]*HTTPTask // task id --> task
	mu    sync.Mutex
}

func NewTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make(map[string]*HTTPTask),
		mu:    sync.Mutex{},
	}
}

func (manager *TaskManager) AddTask(task *HTTPTask) {

	manager.mu.Lock()
	defer manager.mu.Unlock()

	// Stop the old task and start the new task
	if task, ok := manager.tasks[task.ID()]; ok {
		task.stopCh <- struct{}{}
	}

	manager.tasks[task.ID()] = task
	go task.run()
}

func (manager *TaskManager) RemoveTask() error {
	return nil
}

// Get the copy of a task
func (manager *TaskManager) GetTask(id string) (Task, error) {
	return Task{}, nil
}

func (manager *TaskManager) GetTaskPtr() (*Task, error) {
	return nil, nil
}

// List the copy of all task
func (manager *TaskManager) ListTasks() (map[string]HTTPTask, error) {
	return nil, nil
}

func (manager *TaskManager) ListTasksPtr() (map[string]*HTTPTask, error) {
	return nil, nil
}
