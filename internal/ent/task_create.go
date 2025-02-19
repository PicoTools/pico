// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PicoTools/pico/internal/ent/agent"
	"github.com/PicoTools/pico/internal/ent/blobber"
	"github.com/PicoTools/pico/internal/ent/command"
	"github.com/PicoTools/pico/internal/ent/task"
	"github.com/PicoTools/pico/pkg/shared"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetCommandID sets the "command_id" field.
func (tc *TaskCreate) SetCommandID(i int64) *TaskCreate {
	tc.mutation.SetCommandID(i)
	return tc
}

// SetAgentID sets the "agent_id" field.
func (tc *TaskCreate) SetAgentID(u uint32) *TaskCreate {
	tc.mutation.SetAgentID(u)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TaskCreate) SetCreatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetPushedAt sets the "pushed_at" field.
func (tc *TaskCreate) SetPushedAt(t time.Time) *TaskCreate {
	tc.mutation.SetPushedAt(t)
	return tc
}

// SetNillablePushedAt sets the "pushed_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillablePushedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetPushedAt(*t)
	}
	return tc
}

// SetDoneAt sets the "done_at" field.
func (tc *TaskCreate) SetDoneAt(t time.Time) *TaskCreate {
	tc.mutation.SetDoneAt(t)
	return tc
}

// SetNillableDoneAt sets the "done_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableDoneAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetDoneAt(*t)
	}
	return tc
}

// SetStatus sets the "status" field.
func (tc *TaskCreate) SetStatus(ss shared.TaskStatus) *TaskCreate {
	tc.mutation.SetStatus(ss)
	return tc
}

// SetCap sets the "cap" field.
func (tc *TaskCreate) SetCap(s shared.Capability) *TaskCreate {
	tc.mutation.SetCap(s)
	return tc
}

// SetArgsID sets the "args_id" field.
func (tc *TaskCreate) SetArgsID(i int) *TaskCreate {
	tc.mutation.SetArgsID(i)
	return tc
}

// SetOutputID sets the "output_id" field.
func (tc *TaskCreate) SetOutputID(i int) *TaskCreate {
	tc.mutation.SetOutputID(i)
	return tc
}

// SetNillableOutputID sets the "output_id" field if the given value is not nil.
func (tc *TaskCreate) SetNillableOutputID(i *int) *TaskCreate {
	if i != nil {
		tc.SetOutputID(*i)
	}
	return tc
}

// SetOutputBig sets the "output_big" field.
func (tc *TaskCreate) SetOutputBig(b bool) *TaskCreate {
	tc.mutation.SetOutputBig(b)
	return tc
}

// SetNillableOutputBig sets the "output_big" field if the given value is not nil.
func (tc *TaskCreate) SetNillableOutputBig(b *bool) *TaskCreate {
	if b != nil {
		tc.SetOutputBig(*b)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TaskCreate) SetID(i int64) *TaskCreate {
	tc.mutation.SetID(i)
	return tc
}

// SetCommand sets the "command" edge to the Command entity.
func (tc *TaskCreate) SetCommand(c *Command) *TaskCreate {
	return tc.SetCommandID(c.ID)
}

// SetAgent sets the "agent" edge to the Agent entity.
func (tc *TaskCreate) SetAgent(a *Agent) *TaskCreate {
	return tc.SetAgentID(a.ID)
}

// SetBlobberArgsID sets the "blobber_args" edge to the Blobber entity by ID.
func (tc *TaskCreate) SetBlobberArgsID(id int) *TaskCreate {
	tc.mutation.SetBlobberArgsID(id)
	return tc
}

// SetBlobberArgs sets the "blobber_args" edge to the Blobber entity.
func (tc *TaskCreate) SetBlobberArgs(b *Blobber) *TaskCreate {
	return tc.SetBlobberArgsID(b.ID)
}

// SetBlobberOutputID sets the "blobber_output" edge to the Blobber entity by ID.
func (tc *TaskCreate) SetBlobberOutputID(id int) *TaskCreate {
	tc.mutation.SetBlobberOutputID(id)
	return tc
}

// SetNillableBlobberOutputID sets the "blobber_output" edge to the Blobber entity by ID if the given value is not nil.
func (tc *TaskCreate) SetNillableBlobberOutputID(id *int) *TaskCreate {
	if id != nil {
		tc = tc.SetBlobberOutputID(*id)
	}
	return tc
}

// SetBlobberOutput sets the "blobber_output" edge to the Blobber entity.
func (tc *TaskCreate) SetBlobberOutput(b *Blobber) *TaskCreate {
	return tc.SetBlobberOutputID(b.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := task.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.CommandID(); !ok {
		return &ValidationError{Name: "command_id", err: errors.New(`ent: missing required field "Task.command_id"`)}
	}
	if _, ok := tc.mutation.AgentID(); !ok {
		return &ValidationError{Name: "agent_id", err: errors.New(`ent: missing required field "Task.agent_id"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Task.created_at"`)}
	}
	if _, ok := tc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Task.status"`)}
	}
	if v, ok := tc.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Task.status": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Cap(); !ok {
		return &ValidationError{Name: "cap", err: errors.New(`ent: missing required field "Task.cap"`)}
	}
	if v, ok := tc.mutation.Cap(); ok {
		if err := task.CapValidator(v); err != nil {
			return &ValidationError{Name: "cap", err: fmt.Errorf(`ent: validator failed for field "Task.cap": %w`, err)}
		}
	}
	if _, ok := tc.mutation.ArgsID(); !ok {
		return &ValidationError{Name: "args_id", err: errors.New(`ent: missing required field "Task.args_id"`)}
	}
	if len(tc.mutation.CommandIDs()) == 0 {
		return &ValidationError{Name: "command", err: errors.New(`ent: missing required edge "Task.command"`)}
	}
	if len(tc.mutation.AgentIDs()) == 0 {
		return &ValidationError{Name: "agent", err: errors.New(`ent: missing required edge "Task.agent"`)}
	}
	if len(tc.mutation.BlobberArgsIDs()) == 0 {
		return &ValidationError{Name: "blobber_args", err: errors.New(`ent: missing required edge "Task.blobber_args"`)}
	}
	return nil
}

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(task.Table, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt64))
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(task.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.PushedAt(); ok {
		_spec.SetField(task.FieldPushedAt, field.TypeTime, value)
		_node.PushedAt = value
	}
	if value, ok := tc.mutation.DoneAt(); ok {
		_spec.SetField(task.FieldDoneAt, field.TypeTime, value)
		_node.DoneAt = value
	}
	if value, ok := tc.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := tc.mutation.Cap(); ok {
		_spec.SetField(task.FieldCap, field.TypeEnum, value)
		_node.Cap = value
	}
	if value, ok := tc.mutation.OutputBig(); ok {
		_spec.SetField(task.FieldOutputBig, field.TypeBool, value)
		_node.OutputBig = value
	}
	if nodes := tc.mutation.CommandIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.CommandTable,
			Columns: []string{task.CommandColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(command.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CommandID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.AgentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.AgentTable,
			Columns: []string{task.AgentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(agent.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AgentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.BlobberArgsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.BlobberArgsTable,
			Columns: []string{task.BlobberArgsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blobber.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ArgsID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.BlobberOutputIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.BlobberOutputTable,
			Columns: []string{task.BlobberOutputColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blobber.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OutputID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	err      error
	builders []*TaskCreate
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
