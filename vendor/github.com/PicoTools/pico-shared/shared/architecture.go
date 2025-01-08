package shared

import (
	"database/sql/driver"
)

// represents custom type to hold agent's architecture type
type AgentArch uint8

const (
	ArchUnknown AgentArch = iota
	ArchX86
	ArchX64
	ArchArm32
	ArchArm64
)

// Values returns list of strings represented names of arch types
func (AgentArch) Values() []string {
	return []string{
		ArchUnknown.String(),
		ArchX86.String(),
		ArchX64.String(),
		ArchArm32.String(),
		ArchArm64.String(),
	}
}

// Value return database ready value for further processing
func (a AgentArch) Value() (driver.Value, error) {
	return a.String(), nil
}

// String returns string representation of arch types
func (a AgentArch) String() string {
	switch a {
	case ArchX86:
		return "x86"
	case ArchX64:
		return "x64"
	case ArchArm32:
		return "arm32"
	case ArchArm64:
		return "arm64"
	default:
		return "unknown"
	}
}

// Scan converts value to arch type
func (a *AgentArch) Scan(val any) error {
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
	case ArchUnknown.String():
		*a = ArchUnknown
	case ArchX86.String():
		*a = ArchX86
	case ArchX64.String():
		*a = ArchX64
	case ArchArm32.String():
		*a = ArchArm32
	case ArchArm64.String():
		*a = ArchArm64
	}
	return nil
}
