package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hansung080/taskman/task"
)

type ResponseError struct {
	Err error
}

func (e ResponseError) MarshalJSON() ([]byte, error) {
	if e.Err == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%v\"", e.Err)), nil
}

func (e *ResponseError) UnmarshalJSON(data []byte) error {
	var v any
	if err := json.Unmarshal(data, v); err != nil {
		return err
	}
	if v == nil {
		e.Err = nil
		return nil
	}

	switch tv := v.(type) {
	case string:
		if tv == task.ErrTaskNotExist.Error() {
			e.Err = task.ErrTaskNotExist
			return nil
		}
		e.Err = errors.New(tv)
		return nil
	default:
		return errors.New("ResponseError unmarshal failure")
	}
}

type Response struct {
	ID    task.ID       `json:"id,omitempty"`
	Task  task.Task     `json:"task"`
	Error ResponseError `json:"error"`
}
