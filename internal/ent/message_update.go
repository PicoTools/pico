// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PicoTools/pico-shared/shared"
	"github.com/PicoTools/pico/internal/ent/command"
	"github.com/PicoTools/pico/internal/ent/message"
	"github.com/PicoTools/pico/internal/ent/predicate"
)

// MessageUpdate is the builder for updating Message entities.
type MessageUpdate struct {
	config
	hooks    []Hook
	mutation *MessageMutation
}

// Where appends a list predicates to the MessageUpdate builder.
func (mu *MessageUpdate) Where(ps ...predicate.Message) *MessageUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetCommandID sets the "command_id" field.
func (mu *MessageUpdate) SetCommandID(i int64) *MessageUpdate {
	mu.mutation.SetCommandID(i)
	return mu
}

// SetNillableCommandID sets the "command_id" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableCommandID(i *int64) *MessageUpdate {
	if i != nil {
		mu.SetCommandID(*i)
	}
	return mu
}

// SetType sets the "type" field.
func (mu *MessageUpdate) SetType(sm shared.TaskMessage) *MessageUpdate {
	mu.mutation.SetType(sm)
	return mu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableType(sm *shared.TaskMessage) *MessageUpdate {
	if sm != nil {
		mu.SetType(*sm)
	}
	return mu
}

// SetMessage sets the "message" field.
func (mu *MessageUpdate) SetMessage(s string) *MessageUpdate {
	mu.mutation.SetMessage(s)
	return mu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableMessage(s *string) *MessageUpdate {
	if s != nil {
		mu.SetMessage(*s)
	}
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MessageUpdate) SetCreatedAt(t time.Time) *MessageUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mu *MessageUpdate) SetNillableCreatedAt(t *time.Time) *MessageUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// SetCommand sets the "command" edge to the Command entity.
func (mu *MessageUpdate) SetCommand(c *Command) *MessageUpdate {
	return mu.SetCommandID(c.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (mu *MessageUpdate) Mutation() *MessageMutation {
	return mu.mutation
}

// ClearCommand clears the "command" edge to the Command entity.
func (mu *MessageUpdate) ClearCommand() *MessageUpdate {
	mu.mutation.ClearCommand()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MessageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MessageUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MessageUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MessageUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MessageUpdate) check() error {
	if v, ok := mu.mutation.GetType(); ok {
		if err := message.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Message.type": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Message(); ok {
		if err := message.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "Message.message": %w`, err)}
		}
	}
	if mu.mutation.CommandCleared() && len(mu.mutation.CommandIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Message.command"`)
	}
	return nil
}

func (mu *MessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.GetType(); ok {
		_spec.SetField(message.FieldType, field.TypeEnum, value)
	}
	if value, ok := mu.mutation.Message(); ok {
		_spec.SetField(message.FieldMessage, field.TypeString, value)
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.SetField(message.FieldCreatedAt, field.TypeTime, value)
	}
	if mu.mutation.CommandCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.CommandTable,
			Columns: []string{message.CommandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(command.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.CommandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.CommandTable,
			Columns: []string{message.CommandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(command.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MessageUpdateOne is the builder for updating a single Message entity.
type MessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MessageMutation
}

// SetCommandID sets the "command_id" field.
func (muo *MessageUpdateOne) SetCommandID(i int64) *MessageUpdateOne {
	muo.mutation.SetCommandID(i)
	return muo
}

// SetNillableCommandID sets the "command_id" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableCommandID(i *int64) *MessageUpdateOne {
	if i != nil {
		muo.SetCommandID(*i)
	}
	return muo
}

// SetType sets the "type" field.
func (muo *MessageUpdateOne) SetType(sm shared.TaskMessage) *MessageUpdateOne {
	muo.mutation.SetType(sm)
	return muo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableType(sm *shared.TaskMessage) *MessageUpdateOne {
	if sm != nil {
		muo.SetType(*sm)
	}
	return muo
}

// SetMessage sets the "message" field.
func (muo *MessageUpdateOne) SetMessage(s string) *MessageUpdateOne {
	muo.mutation.SetMessage(s)
	return muo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableMessage(s *string) *MessageUpdateOne {
	if s != nil {
		muo.SetMessage(*s)
	}
	return muo
}

// SetCreatedAt sets the "created_at" field.
func (muo *MessageUpdateOne) SetCreatedAt(t time.Time) *MessageUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (muo *MessageUpdateOne) SetNillableCreatedAt(t *time.Time) *MessageUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// SetCommand sets the "command" edge to the Command entity.
func (muo *MessageUpdateOne) SetCommand(c *Command) *MessageUpdateOne {
	return muo.SetCommandID(c.ID)
}

// Mutation returns the MessageMutation object of the builder.
func (muo *MessageUpdateOne) Mutation() *MessageMutation {
	return muo.mutation
}

// ClearCommand clears the "command" edge to the Command entity.
func (muo *MessageUpdateOne) ClearCommand() *MessageUpdateOne {
	muo.mutation.ClearCommand()
	return muo
}

// Where appends a list predicates to the MessageUpdate builder.
func (muo *MessageUpdateOne) Where(ps ...predicate.Message) *MessageUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MessageUpdateOne) Select(field string, fields ...string) *MessageUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Message entity.
func (muo *MessageUpdateOne) Save(ctx context.Context) (*Message, error) {
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MessageUpdateOne) SaveX(ctx context.Context) *Message {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MessageUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MessageUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MessageUpdateOne) check() error {
	if v, ok := muo.mutation.GetType(); ok {
		if err := message.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Message.type": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Message(); ok {
		if err := message.MessageValidator(v); err != nil {
			return &ValidationError{Name: "message", err: fmt.Errorf(`ent: validator failed for field "Message.message": %w`, err)}
		}
	}
	if muo.mutation.CommandCleared() && len(muo.mutation.CommandIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Message.command"`)
	}
	return nil
}

func (muo *MessageUpdateOne) sqlSave(ctx context.Context) (_node *Message, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(message.Table, message.Columns, sqlgraph.NewFieldSpec(message.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Message.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, message.FieldID)
		for _, f := range fields {
			if !message.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != message.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.GetType(); ok {
		_spec.SetField(message.FieldType, field.TypeEnum, value)
	}
	if value, ok := muo.mutation.Message(); ok {
		_spec.SetField(message.FieldMessage, field.TypeString, value)
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.SetField(message.FieldCreatedAt, field.TypeTime, value)
	}
	if muo.mutation.CommandCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.CommandTable,
			Columns: []string{message.CommandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(command.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.CommandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   message.CommandTable,
			Columns: []string{message.CommandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(command.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Message{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{message.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}