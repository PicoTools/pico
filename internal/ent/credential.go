// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/PicoTools/pico/internal/ent/credential"
)

// Credential is the model entity for the Credential schema.
type Credential struct {
	config `json:"-"`
	// ID of the ent.
	// credential ID
	ID int64 `json:"id,omitempty"`
	// Time when entity was created
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Time when entity was updated
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Time when entity was soft-deleted
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// username of credential
	Username string `json:"username,omitempty"`
	// password/hash/secret
	Secret string `json:"secret,omitempty"`
	// realm of host (or zone where credentials are valid)
	Realm string `json:"realm,omitempty"`
	// host from which creds has been extracted
	Host string `json:"host,omitempty"`
	// note of credential
	Note string `json:"note,omitempty"`
	// color of entity
	Color        uint32 `json:"color,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Credential) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case credential.FieldID, credential.FieldColor:
			values[i] = new(sql.NullInt64)
		case credential.FieldUsername, credential.FieldSecret, credential.FieldRealm, credential.FieldHost, credential.FieldNote:
			values[i] = new(sql.NullString)
		case credential.FieldCreatedAt, credential.FieldUpdatedAt, credential.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Credential fields.
func (c *Credential) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case credential.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int64(value.Int64)
		case credential.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case credential.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case credential.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				c.DeletedAt = value.Time
			}
		case credential.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				c.Username = value.String
			}
		case credential.FieldSecret:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value.Valid {
				c.Secret = value.String
			}
		case credential.FieldRealm:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field realm", values[i])
			} else if value.Valid {
				c.Realm = value.String
			}
		case credential.FieldHost:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field host", values[i])
			} else if value.Valid {
				c.Host = value.String
			}
		case credential.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				c.Note = value.String
			}
		case credential.FieldColor:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				c.Color = uint32(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Credential.
// This includes values selected through modifiers, order, etc.
func (c *Credential) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Credential.
// Note that you need to call Credential.Unwrap() before calling this method if this Credential
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Credential) Update() *CredentialUpdateOne {
	return NewCredentialClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Credential entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Credential) Unwrap() *Credential {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Credential is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Credential) String() string {
	var builder strings.Builder
	builder.WriteString("Credential(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(c.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(c.Username)
	builder.WriteString(", ")
	builder.WriteString("secret=")
	builder.WriteString(c.Secret)
	builder.WriteString(", ")
	builder.WriteString("realm=")
	builder.WriteString(c.Realm)
	builder.WriteString(", ")
	builder.WriteString("host=")
	builder.WriteString(c.Host)
	builder.WriteString(", ")
	builder.WriteString("note=")
	builder.WriteString(c.Note)
	builder.WriteString(", ")
	builder.WriteString("color=")
	builder.WriteString(fmt.Sprintf("%v", c.Color))
	builder.WriteByte(')')
	return builder.String()
}

// Credentials is a parsable slice of Credential.
type Credentials []*Credential
