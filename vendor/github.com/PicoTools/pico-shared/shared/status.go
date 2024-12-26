package shared

import (
	"database/sql/driver"
)

// represents custom type to hold task's status
type TaskStatus uint8

const (
	StatusNew TaskStatus = iota
	StatusInProgress
	StatusCancelled
	StatusSuccess
	StatusError
)

// Values returns list of strings represented names of task's statuses
func (TaskStatus) Values() []string {
	return []string{
		StatusNew.String(),
		StatusInProgress.String(),
		StatusCancelled.String(),
		StatusSuccess.String(),
		StatusError.String(),
	}
}

// Value return database ready value for further processing
func (t TaskStatus) Value() (driver.Value, error) {
	return t.String(), nil
}

// String returns string representation of task's status
func (t TaskStatus) String() string {
	switch t {
	case StatusNew:
		return "new"
	case StatusInProgress:
		return "in-progress"
	case StatusCancelled:
		return "cancelled"
	case StatusSuccess:
		return "success"
	case StatusError:
		return "error"
	default:
		return "unknown"
	}
}

// Scan converts value to task status
func (t *TaskStatus) Scan(val any) error {
	var s string

	switch v := val.(type) {
	case nil:
		return nil
	case []uint8:
		s = string(v)
	case string:
		s = v
	}

	switch s {
	case StatusNew.String():
		*t = StatusNew
	case StatusInProgress.String():
		*t = StatusInProgress
	case StatusCancelled.String():
		*t = StatusCancelled
	case StatusSuccess.String():
		*t = StatusSuccess
	case StatusError.String():
		*t = StatusError
	}
	return nil
}
