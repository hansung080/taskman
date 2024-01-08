package task

import "errors"

var ErrTaskNotExist = errors.New("task does not exist")

type ID string

type Writer interface {
	Add(t Task) (ID, error)
	Update(id ID, t Task) error
	Delete(id ID) error
}

type Reader interface {
	Get(id ID) (Task, error)
}

type Accessor interface {
	Writer
	Reader
}
