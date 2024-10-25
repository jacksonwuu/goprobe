package task

import "sync"

// Manager manage a branch of probe task,
// it will start do probe for these probe jobs.
type Manager struct {
	tasks map[string]Task
	mu    sync.Mutex
}

func NewManager() Manager {
	return Manager{
		tasks: map[string]Task{},
		mu:    sync.Mutex{},
	}
}

func (manager *Manager) Run() {
}

func (manager *Manager) AddTask(task *Task) error {
	return nil
}

func (manager *Manager) RemoveTask() error {
	return nil
}
