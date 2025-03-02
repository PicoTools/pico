// Code generated by ent, DO NOT EDIT.

package command

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/PicoTools/pico/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.Command {
	return predicate.Command(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.Command {
	return predicate.Command(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.Command {
	return predicate.Command(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.Command {
	return predicate.Command(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.Command {
	return predicate.Command(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.Command {
	return predicate.Command(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.Command {
	return predicate.Command(sql.FieldLTE(FieldID, id))
}

// AgentID applies equality check predicate on the "agent_id" field. It's identical to AgentIDEQ.
func AgentID(v uint32) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldAgentID, v))
}

// Cmd applies equality check predicate on the "cmd" field. It's identical to CmdEQ.
func Cmd(v string) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldCmd, v))
}

// Visible applies equality check predicate on the "visible" field. It's identical to VisibleEQ.
func Visible(v bool) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldVisible, v))
}

// AuthorID applies equality check predicate on the "author_id" field. It's identical to AuthorIDEQ.
func AuthorID(v int64) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldAuthorID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldCreatedAt, v))
}

// ClosedAt applies equality check predicate on the "closed_at" field. It's identical to ClosedAtEQ.
func ClosedAt(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldClosedAt, v))
}

// AgentIDEQ applies the EQ predicate on the "agent_id" field.
func AgentIDEQ(v uint32) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldAgentID, v))
}

// AgentIDNEQ applies the NEQ predicate on the "agent_id" field.
func AgentIDNEQ(v uint32) predicate.Command {
	return predicate.Command(sql.FieldNEQ(FieldAgentID, v))
}

// AgentIDIn applies the In predicate on the "agent_id" field.
func AgentIDIn(vs ...uint32) predicate.Command {
	return predicate.Command(sql.FieldIn(FieldAgentID, vs...))
}

// AgentIDNotIn applies the NotIn predicate on the "agent_id" field.
func AgentIDNotIn(vs ...uint32) predicate.Command {
	return predicate.Command(sql.FieldNotIn(FieldAgentID, vs...))
}

// CmdEQ applies the EQ predicate on the "cmd" field.
func CmdEQ(v string) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldCmd, v))
}

// CmdNEQ applies the NEQ predicate on the "cmd" field.
func CmdNEQ(v string) predicate.Command {
	return predicate.Command(sql.FieldNEQ(FieldCmd, v))
}

// CmdIn applies the In predicate on the "cmd" field.
func CmdIn(vs ...string) predicate.Command {
	return predicate.Command(sql.FieldIn(FieldCmd, vs...))
}

// CmdNotIn applies the NotIn predicate on the "cmd" field.
func CmdNotIn(vs ...string) predicate.Command {
	return predicate.Command(sql.FieldNotIn(FieldCmd, vs...))
}

// CmdGT applies the GT predicate on the "cmd" field.
func CmdGT(v string) predicate.Command {
	return predicate.Command(sql.FieldGT(FieldCmd, v))
}

// CmdGTE applies the GTE predicate on the "cmd" field.
func CmdGTE(v string) predicate.Command {
	return predicate.Command(sql.FieldGTE(FieldCmd, v))
}

// CmdLT applies the LT predicate on the "cmd" field.
func CmdLT(v string) predicate.Command {
	return predicate.Command(sql.FieldLT(FieldCmd, v))
}

// CmdLTE applies the LTE predicate on the "cmd" field.
func CmdLTE(v string) predicate.Command {
	return predicate.Command(sql.FieldLTE(FieldCmd, v))
}

// CmdContains applies the Contains predicate on the "cmd" field.
func CmdContains(v string) predicate.Command {
	return predicate.Command(sql.FieldContains(FieldCmd, v))
}

// CmdHasPrefix applies the HasPrefix predicate on the "cmd" field.
func CmdHasPrefix(v string) predicate.Command {
	return predicate.Command(sql.FieldHasPrefix(FieldCmd, v))
}

// CmdHasSuffix applies the HasSuffix predicate on the "cmd" field.
func CmdHasSuffix(v string) predicate.Command {
	return predicate.Command(sql.FieldHasSuffix(FieldCmd, v))
}

// CmdEqualFold applies the EqualFold predicate on the "cmd" field.
func CmdEqualFold(v string) predicate.Command {
	return predicate.Command(sql.FieldEqualFold(FieldCmd, v))
}

// CmdContainsFold applies the ContainsFold predicate on the "cmd" field.
func CmdContainsFold(v string) predicate.Command {
	return predicate.Command(sql.FieldContainsFold(FieldCmd, v))
}

// VisibleEQ applies the EQ predicate on the "visible" field.
func VisibleEQ(v bool) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldVisible, v))
}

// VisibleNEQ applies the NEQ predicate on the "visible" field.
func VisibleNEQ(v bool) predicate.Command {
	return predicate.Command(sql.FieldNEQ(FieldVisible, v))
}

// AuthorIDEQ applies the EQ predicate on the "author_id" field.
func AuthorIDEQ(v int64) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldAuthorID, v))
}

// AuthorIDNEQ applies the NEQ predicate on the "author_id" field.
func AuthorIDNEQ(v int64) predicate.Command {
	return predicate.Command(sql.FieldNEQ(FieldAuthorID, v))
}

// AuthorIDIn applies the In predicate on the "author_id" field.
func AuthorIDIn(vs ...int64) predicate.Command {
	return predicate.Command(sql.FieldIn(FieldAuthorID, vs...))
}

// AuthorIDNotIn applies the NotIn predicate on the "author_id" field.
func AuthorIDNotIn(vs ...int64) predicate.Command {
	return predicate.Command(sql.FieldNotIn(FieldAuthorID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Command {
	return predicate.Command(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Command {
	return predicate.Command(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldLTE(FieldCreatedAt, v))
}

// ClosedAtEQ applies the EQ predicate on the "closed_at" field.
func ClosedAtEQ(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldEQ(FieldClosedAt, v))
}

// ClosedAtNEQ applies the NEQ predicate on the "closed_at" field.
func ClosedAtNEQ(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldNEQ(FieldClosedAt, v))
}

// ClosedAtIn applies the In predicate on the "closed_at" field.
func ClosedAtIn(vs ...time.Time) predicate.Command {
	return predicate.Command(sql.FieldIn(FieldClosedAt, vs...))
}

// ClosedAtNotIn applies the NotIn predicate on the "closed_at" field.
func ClosedAtNotIn(vs ...time.Time) predicate.Command {
	return predicate.Command(sql.FieldNotIn(FieldClosedAt, vs...))
}

// ClosedAtGT applies the GT predicate on the "closed_at" field.
func ClosedAtGT(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldGT(FieldClosedAt, v))
}

// ClosedAtGTE applies the GTE predicate on the "closed_at" field.
func ClosedAtGTE(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldGTE(FieldClosedAt, v))
}

// ClosedAtLT applies the LT predicate on the "closed_at" field.
func ClosedAtLT(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldLT(FieldClosedAt, v))
}

// ClosedAtLTE applies the LTE predicate on the "closed_at" field.
func ClosedAtLTE(v time.Time) predicate.Command {
	return predicate.Command(sql.FieldLTE(FieldClosedAt, v))
}

// ClosedAtIsNil applies the IsNil predicate on the "closed_at" field.
func ClosedAtIsNil() predicate.Command {
	return predicate.Command(sql.FieldIsNull(FieldClosedAt))
}

// ClosedAtNotNil applies the NotNil predicate on the "closed_at" field.
func ClosedAtNotNil() predicate.Command {
	return predicate.Command(sql.FieldNotNull(FieldClosedAt))
}

// HasAgent applies the HasEdge predicate on the "agent" edge.
func HasAgent() predicate.Command {
	return predicate.Command(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AgentTable, AgentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAgentWith applies the HasEdge predicate on the "agent" edge with a given conditions (other predicates).
func HasAgentWith(preds ...predicate.Agent) predicate.Command {
	return predicate.Command(func(s *sql.Selector) {
		step := newAgentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOperator applies the HasEdge predicate on the "operator" edge.
func HasOperator() predicate.Command {
	return predicate.Command(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OperatorTable, OperatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOperatorWith applies the HasEdge predicate on the "operator" edge with a given conditions (other predicates).
func HasOperatorWith(preds ...predicate.Operator) predicate.Command {
	return predicate.Command(func(s *sql.Selector) {
		step := newOperatorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMessage applies the HasEdge predicate on the "message" edge.
func HasMessage() predicate.Command {
	return predicate.Command(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MessageTable, MessageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMessageWith applies the HasEdge predicate on the "message" edge with a given conditions (other predicates).
func HasMessageWith(preds ...predicate.Message) predicate.Command {
	return predicate.Command(func(s *sql.Selector) {
		step := newMessageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTask applies the HasEdge predicate on the "task" edge.
func HasTask() predicate.Command {
	return predicate.Command(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TaskTable, TaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskWith applies the HasEdge predicate on the "task" edge with a given conditions (other predicates).
func HasTaskWith(preds ...predicate.Task) predicate.Command {
	return predicate.Command(func(s *sql.Selector) {
		step := newTaskStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Command) predicate.Command {
	return predicate.Command(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Command) predicate.Command {
	return predicate.Command(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Command) predicate.Command {
	return predicate.Command(sql.NotPredicates(p))
}
