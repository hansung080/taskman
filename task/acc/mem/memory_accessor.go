package mem

import (
	"fmt"

	"github.com/hansung080/taskman/task"
)

type MemoryAccessor struct {
	tasks  map[task.ID]task.Task
	nextID int64
}

func NewMemoryAccessor() *MemoryAccessor {
	return &MemoryAccessor{
		tasks:  map[task.ID]task.Task{},
		nextID: int64(1),
	}
}

func (m *MemoryAccessor) Add(t task.Task) (task.ID, error) {
	id := task.ID(fmt.Sprint(m.nextID))
	m.nextID++
	m.tasks[id] = t
	return id, nil
}

func (m *MemoryAccessor) Update(id task.ID, t task.Task) error {
	if _, exists := m.tasks[id]; !exists {
		return task.ErrTaskNotExist
	}
	m.tasks[id] = t
	return nil
}

func (m *MemoryAccessor) Delete(id task.ID) error {
	if _, exists := m.tasks[id]; !exists {
		return task.ErrTaskNotExist
	}
	delete(m.tasks, id)
	return nil
}

func (m *MemoryAccessor) Get(id task.ID) (task.Task, error) {
	t, exists := m.tasks[id]
	if !exists {
		return task.Task{}, task.ErrTaskNotExist
	}
	return t, nil
}
