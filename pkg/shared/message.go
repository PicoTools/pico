package shared

import (
	"database/sql/driver"
)

// represents custom type to hold task messages
type TaskMessage uint8

const (
	NotifyMessage TaskMessage = iota
	InfoMessage
	WarningMessage
	ErrorMessage
)

// Values returns list of strings represented names of task messages
func (t TaskMessage) Values() []string {
	return []string{
		NotifyMessage.String(),
		InfoMessage.String(),
		WarningMessage.String(),
		ErrorMessage.String(),
	}
}

// Value return database ready value for further processing
func (t TaskMessage) Value() (driver.Value, error) {
	return t.String(), nil
}

// String returns string representation of task messages types
func (t TaskMessage) String() string {
	switch t {
	case NotifyMessage:
		return "notify"
	case InfoMessage:
		return "info"
	case WarningMessage:
		return "warning"
	case ErrorMessage:
		return "error"
	default:
		return "unknown"
	}
}

// Scan converts value to capability
func (t *TaskMessage) Scan(val any) error {
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
	case NotifyMessage.String():
		*t = NotifyMessage
	case InfoMessage.String():
		*t = InfoMessage
	case WarningMessage.String():
		*t = WarningMessage
	case ErrorMessage.String():
		*t = ErrorMessage
	}
	return nil
}
