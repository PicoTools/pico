package shared

import (
	"database/sql/driver"
)

// represents custom type to hold ant's operating system type
type AntOs uint8

const (
	OsUnknown AntOs = iota
	OsLinux
	OsWindows
	OsMac
)

// Values returns list of strings represented names of capability types
func (AntOs) Values() []string {
	return []string{
		OsUnknown.String(),
		OsLinux.String(),
		OsWindows.String(),
		OsMac.String(),
	}
}

// Value return database ready value for further processing
func (a AntOs) Value() (driver.Value, error) {
	return a.String(), nil
}

// Scan converts value to operating system type
func (a *AntOs) Scan(val any) error {
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
	case OsUnknown.String():
		*a = OsUnknown
	case OsLinux.String():
		*a = OsLinux
	case OsWindows.String():
		*a = OsWindows
	case OsMac.String():
		*a = OsMac
	}
	return nil
}

// String returns string representation of operating system type
func (a AntOs) String() string {
	switch a {
	case OsLinux:
		return "linux"
	case OsWindows:
		return "windows"
	case OsMac:
		return "macos"
	default:
		return "unknown"
	}
}

// String returns short string representation of operating system type
func (a AntOs) StringShort() string {
	switch a {
	case OsLinux:
		return "lin"
	case OsWindows:
		return "win"
	case OsMac:
		return "mac"
	default:
		return "???"
	}
}
