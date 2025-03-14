// Code generated by ent, DO NOT EDIT.

package message

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/PicoTools/pico/internal/ent/predicate"
	"github.com/PicoTools/pico/pkg/shared"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldID, id))
}

// CommandID applies equality check predicate on the "command_id" field. It's identical to CommandIDEQ.
func CommandID(v int64) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldCommandID, v))
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldMessage, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldCreatedAt, v))
}

// CommandIDEQ applies the EQ predicate on the "command_id" field.
func CommandIDEQ(v int64) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldCommandID, v))
}

// CommandIDNEQ applies the NEQ predicate on the "command_id" field.
func CommandIDNEQ(v int64) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldCommandID, v))
}

// CommandIDIn applies the In predicate on the "command_id" field.
func CommandIDIn(vs ...int64) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldCommandID, vs...))
}

// CommandIDNotIn applies the NotIn predicate on the "command_id" field.
func CommandIDNotIn(vs ...int64) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldCommandID, vs...))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v shared.TaskMessage) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v shared.TaskMessage) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...shared.TaskMessage) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...shared.TaskMessage) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldType, vs...))
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldMessage, v))
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldMessage, v))
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldMessage, vs...))
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldMessage, vs...))
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldMessage, v))
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldMessage, v))
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldMessage, v))
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldMessage, v))
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.Message {
	return predicate.Message(sql.FieldContains(FieldMessage, v))
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasPrefix(FieldMessage, v))
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasSuffix(FieldMessage, v))
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.Message {
	return predicate.Message(sql.FieldEqualFold(FieldMessage, v))
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.Message {
	return predicate.Message(sql.FieldContainsFold(FieldMessage, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldCreatedAt, v))
}

// HasCommand applies the HasEdge predicate on the "command" edge.
func HasCommand() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CommandTable, CommandColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommandWith applies the HasEdge predicate on the "command" edge with a given conditions (other predicates).
func HasCommandWith(preds ...predicate.Command) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := newCommandStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Message) predicate.Message {
	return predicate.Message(sql.NotPredicates(p))
}
