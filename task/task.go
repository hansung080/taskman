package task

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type Status int

const (
	Unknown Status = iota
	Todo
	Done
)

func (s Status) String() string {
	switch s {
	case Unknown:
		return "Unknown"
	case Todo:
		return "Todo"
	case Done:
		return "Done"
	default:
		return ""
	}
}

func (s Status) MarshalJSON() ([]byte, error) {
	str := s.String()
	if str == "" {
		return nil, errors.New("empty status")
	}
	return []byte(fmt.Sprintf("\"%s\"", str)), nil
}

func (s *Status) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"Unknown"`:
		*s = Unknown
	case `"Todo"`:
		*s = Todo
	case `"Done"`:
		*s = Done
	default:
		return fmt.Errorf("invalid status: %s", string(data))
	}
	return nil
}

type Deadline struct {
	time.Time
}

func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

func (d Deadline) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, d.Unix(), 10), nil
}

func (d *Deadline) UnmarshalJSON(data []byte) error {
	unix, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.Unix(unix, 0)
	return nil
}

type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   Status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subTasks,omitempty"`
}

func (t Task) String() string {
	check := " "
	if t.Status == Done {
		check = "v"
	}
	return fmt.Sprintf("[%s] %s, %s, %d", check, t.Title, t.Deadline, t.Priority)
}

type IncludeSubTasks Task

func (t IncludeSubTasks) indentedString(indent string) string {
	str := indent + Task(t).String()
	for _, st := range t.SubTasks {
		str += "\n" + IncludeSubTasks(st).indentedString(indent+"    ")
	}
	return str
}

func (t IncludeSubTasks) String() string {
	return t.indentedString("")
}
