// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/PicoTools/pico/internal/ent/operator"
)

// Operator is the model entity for the Operator schema.
type Operator struct {
	config `json:"-"`
	// ID of the ent.
	// operator ID
	ID int64 `json:"id,omitempty"`
	// Time when entity was created
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Time when entity was updated
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Time when entity was soft-deleted
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// username of operator
	Username string `json:"username,omitempty"`
	// access token for operator
	Token string `json:"token,omitempty"`
	// color of entity
	Color uint32 `json:"color,omitempty"`
	// last activity of operator
	Last time.Time `json:"last,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OperatorQuery when eager-loading is set.
	Edges        OperatorEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OperatorEdges holds the relations/edges for other nodes in the graph.
type OperatorEdges struct {
	// Chat holds the value of the chat edge.
	Chat []*Chat `json:"chat,omitempty"`
	// Command holds the value of the command edge.
	Command []*Command `json:"command,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ChatOrErr returns the Chat value or an error if the edge
// was not loaded in eager-loading.
func (e OperatorEdges) ChatOrErr() ([]*Chat, error) {
	if e.loadedTypes[0] {
		return e.Chat, nil
	}
	return nil, &NotLoadedError{edge: "chat"}
}

// CommandOrErr returns the Command value or an error if the edge
// was not loaded in eager-loading.
func (e OperatorEdges) CommandOrErr() ([]*Command, error) {
	if e.loadedTypes[1] {
		return e.Command, nil
	}
	return nil, &NotLoadedError{edge: "command"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Operator) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case operator.FieldID, operator.FieldColor:
			values[i] = new(sql.NullInt64)
		case operator.FieldUsername, operator.FieldToken:
			values[i] = new(sql.NullString)
		case operator.FieldCreatedAt, operator.FieldUpdatedAt, operator.FieldDeletedAt, operator.FieldLast:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Operator fields.
func (o *Operator) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case operator.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int64(value.Int64)
		case operator.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case operator.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case operator.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				o.DeletedAt = value.Time
			}
		case operator.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				o.Username = value.String
			}
		case operator.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				o.Token = value.String
			}
		case operator.FieldColor:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				o.Color = uint32(value.Int64)
			}
		case operator.FieldLast:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last", values[i])
			} else if value.Valid {
				o.Last = value.Time
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Operator.
// This includes values selected through modifiers, order, etc.
func (o *Operator) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryChat queries the "chat" edge of the Operator entity.
func (o *Operator) QueryChat() *ChatQuery {
	return NewOperatorClient(o.config).QueryChat(o)
}

// QueryCommand queries the "command" edge of the Operator entity.
func (o *Operator) QueryCommand() *CommandQuery {
	return NewOperatorClient(o.config).QueryCommand(o)
}

// Update returns a builder for updating this Operator.
// Note that you need to call Operator.Unwrap() before calling this method if this Operator
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Operator) Update() *OperatorUpdateOne {
	return NewOperatorClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Operator entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Operator) Unwrap() *Operator {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Operator is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Operator) String() string {
	var builder strings.Builder
	builder.WriteString("Operator(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(o.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(o.Username)
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(o.Token)
	builder.WriteString(", ")
	builder.WriteString("color=")
	builder.WriteString(fmt.Sprintf("%v", o.Color))
	builder.WriteString(", ")
	builder.WriteString("last=")
	builder.WriteString(o.Last.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Operators is a parsable slice of Operator.
type Operators []*Operator
