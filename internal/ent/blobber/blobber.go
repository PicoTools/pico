// Code generated by ent, DO NOT EDIT.

package blobber

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the blobber type in the database.
	Label = "blobber"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldHash holds the string denoting the hash field in the database.
	FieldHash = "hash"
	// FieldBlob holds the string denoting the blob field in the database.
	FieldBlob = "blob"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// EdgeTaskArgs holds the string denoting the task_args edge name in mutations.
	EdgeTaskArgs = "task_args"
	// EdgeTaskOutput holds the string denoting the task_output edge name in mutations.
	EdgeTaskOutput = "task_output"
	// Table holds the table name of the blobber in the database.
	Table = "blobber"
	// TaskArgsTable is the table that holds the task_args relation/edge.
	TaskArgsTable = "task"
	// TaskArgsInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskArgsInverseTable = "task"
	// TaskArgsColumn is the table column denoting the task_args relation/edge.
	TaskArgsColumn = "args_id"
	// TaskOutputTable is the table that holds the task_output relation/edge.
	TaskOutputTable = "task"
	// TaskOutputInverseTable is the table name for the Task entity.
	// It exists in this package in order to avoid circular dependency with the "task" package.
	TaskOutputInverseTable = "task"
	// TaskOutputColumn is the table column denoting the task_output relation/edge.
	TaskOutputColumn = "output_id"
)

// Columns holds all SQL columns for blobber fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldHash,
	FieldBlob,
	FieldSize,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/PicoTools/pico/internal/ent/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
)

// OrderOption defines the ordering options for the Blobber queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// ByTaskArgsCount orders the results by task_args count.
func ByTaskArgsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTaskArgsStep(), opts...)
	}
}

// ByTaskArgs orders the results by task_args terms.
func ByTaskArgs(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTaskArgsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTaskOutputCount orders the results by task_output count.
func ByTaskOutputCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTaskOutputStep(), opts...)
	}
}

// ByTaskOutput orders the results by task_output terms.
func ByTaskOutput(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTaskOutputStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newTaskArgsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TaskArgsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TaskArgsTable, TaskArgsColumn),
	)
}
func newTaskOutputStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TaskOutputInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TaskOutputTable, TaskOutputColumn),
	)
}