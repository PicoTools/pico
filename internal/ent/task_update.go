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
	"github.com/PicoTools/pico/internal/ent/ant"
	"github.com/PicoTools/pico/internal/ent/blobber"
	"github.com/PicoTools/pico/internal/ent/command"
	"github.com/PicoTools/pico/internal/ent/predicate"
	"github.com/PicoTools/pico/internal/ent/task"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetCommandID sets the "command_id" field.
func (tu *TaskUpdate) SetCommandID(i int64) *TaskUpdate {
	tu.mutation.SetCommandID(i)
	return tu
}

// SetNillableCommandID sets the "command_id" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCommandID(i *int64) *TaskUpdate {
	if i != nil {
		tu.SetCommandID(*i)
	}
	return tu
}

// SetAntID sets the "ant_id" field.
func (tu *TaskUpdate) SetAntID(u uint32) *TaskUpdate {
	tu.mutation.SetAntID(u)
	return tu
}

// SetNillableAntID sets the "ant_id" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableAntID(u *uint32) *TaskUpdate {
	if u != nil {
		tu.SetAntID(*u)
	}
	return tu
}

// SetCreatedAt sets the "created_at" field.
func (tu *TaskUpdate) SetCreatedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCreatedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetPushedAt sets the "pushed_at" field.
func (tu *TaskUpdate) SetPushedAt(t time.Time) *TaskUpdate {
	tu.mutation.SetPushedAt(t)
	return tu
}

// SetNillablePushedAt sets the "pushed_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillablePushedAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetPushedAt(*t)
	}
	return tu
}

// ClearPushedAt clears the value of the "pushed_at" field.
func (tu *TaskUpdate) ClearPushedAt() *TaskUpdate {
	tu.mutation.ClearPushedAt()
	return tu
}

// SetDoneAt sets the "done_at" field.
func (tu *TaskUpdate) SetDoneAt(t time.Time) *TaskUpdate {
	tu.mutation.SetDoneAt(t)
	return tu
}

// SetNillableDoneAt sets the "done_at" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDoneAt(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetDoneAt(*t)
	}
	return tu
}

// ClearDoneAt clears the value of the "done_at" field.
func (tu *TaskUpdate) ClearDoneAt() *TaskUpdate {
	tu.mutation.ClearDoneAt()
	return tu
}

// SetStatus sets the "status" field.
func (tu *TaskUpdate) SetStatus(ss shared.TaskStatus) *TaskUpdate {
	tu.mutation.SetStatus(ss)
	return tu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableStatus(ss *shared.TaskStatus) *TaskUpdate {
	if ss != nil {
		tu.SetStatus(*ss)
	}
	return tu
}

// SetCap sets the "cap" field.
func (tu *TaskUpdate) SetCap(s shared.Capability) *TaskUpdate {
	tu.mutation.SetCap(s)
	return tu
}

// SetNillableCap sets the "cap" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableCap(s *shared.Capability) *TaskUpdate {
	if s != nil {
		tu.SetCap(*s)
	}
	return tu
}

// SetArgsID sets the "args_id" field.
func (tu *TaskUpdate) SetArgsID(i int) *TaskUpdate {
	tu.mutation.SetArgsID(i)
	return tu
}

// SetNillableArgsID sets the "args_id" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableArgsID(i *int) *TaskUpdate {
	if i != nil {
		tu.SetArgsID(*i)
	}
	return tu
}

// SetOutputID sets the "output_id" field.
func (tu *TaskUpdate) SetOutputID(i int) *TaskUpdate {
	tu.mutation.SetOutputID(i)
	return tu
}

// SetNillableOutputID sets the "output_id" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOutputID(i *int) *TaskUpdate {
	if i != nil {
		tu.SetOutputID(*i)
	}
	return tu
}

// ClearOutputID clears the value of the "output_id" field.
func (tu *TaskUpdate) ClearOutputID() *TaskUpdate {
	tu.mutation.ClearOutputID()
	return tu
}

// SetOutputBig sets the "output_big" field.
func (tu *TaskUpdate) SetOutputBig(b bool) *TaskUpdate {
	tu.mutation.SetOutputBig(b)
	return tu
}

// SetNillableOutputBig sets the "output_big" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOutputBig(b *bool) *TaskUpdate {
	if b != nil {
		tu.SetOutputBig(*b)
	}
	return tu
}

// ClearOutputBig clears the value of the "output_big" field.
func (tu *TaskUpdate) ClearOutputBig() *TaskUpdate {
	tu.mutation.ClearOutputBig()
	return tu
}

// SetCommand sets the "command" edge to the Command entity.
func (tu *TaskUpdate) SetCommand(c *Command) *TaskUpdate {
	return tu.SetCommandID(c.ID)
}

// SetAnt sets the "ant" edge to the Ant entity.
func (tu *TaskUpdate) SetAnt(a *Ant) *TaskUpdate {
	return tu.SetAntID(a.ID)
}

// SetBlobberArgsID sets the "blobber_args" edge to the Blobber entity by ID.
func (tu *TaskUpdate) SetBlobberArgsID(id int) *TaskUpdate {
	tu.mutation.SetBlobberArgsID(id)
	return tu
}

// SetBlobberArgs sets the "blobber_args" edge to the Blobber entity.
func (tu *TaskUpdate) SetBlobberArgs(b *Blobber) *TaskUpdate {
	return tu.SetBlobberArgsID(b.ID)
}

// SetBlobberOutputID sets the "blobber_output" edge to the Blobber entity by ID.
func (tu *TaskUpdate) SetBlobberOutputID(id int) *TaskUpdate {
	tu.mutation.SetBlobberOutputID(id)
	return tu
}

// SetNillableBlobberOutputID sets the "blobber_output" edge to the Blobber entity by ID if the given value is not nil.
func (tu *TaskUpdate) SetNillableBlobberOutputID(id *int) *TaskUpdate {
	if id != nil {
		tu = tu.SetBlobberOutputID(*id)
	}
	return tu
}

// SetBlobberOutput sets the "blobber_output" edge to the Blobber entity.
func (tu *TaskUpdate) SetBlobberOutput(b *Blobber) *TaskUpdate {
	return tu.SetBlobberOutputID(b.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// ClearCommand clears the "command" edge to the Command entity.
func (tu *TaskUpdate) ClearCommand() *TaskUpdate {
	tu.mutation.ClearCommand()
	return tu
}

// ClearAnt clears the "ant" edge to the Ant entity.
func (tu *TaskUpdate) ClearAnt() *TaskUpdate {
	tu.mutation.ClearAnt()
	return tu
}

// ClearBlobberArgs clears the "blobber_args" edge to the Blobber entity.
func (tu *TaskUpdate) ClearBlobberArgs() *TaskUpdate {
	tu.mutation.ClearBlobberArgs()
	return tu
}

// ClearBlobberOutput clears the "blobber_output" edge to the Blobber entity.
func (tu *TaskUpdate) ClearBlobberOutput() *TaskUpdate {
	tu.mutation.ClearBlobberOutput()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Task.status": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Cap(); ok {
		if err := task.CapValidator(v); err != nil {
			return &ValidationError{Name: "cap", err: fmt.Errorf(`ent: validator failed for field "Task.cap": %w`, err)}
		}
	}
	if tu.mutation.CommandCleared() && len(tu.mutation.CommandIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Task.command"`)
	}
	if tu.mutation.AntCleared() && len(tu.mutation.AntIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Task.ant"`)
	}
	if tu.mutation.BlobberArgsCleared() && len(tu.mutation.BlobberArgsIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Task.blobber_args"`)
	}
	return nil
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt64))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(task.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.PushedAt(); ok {
		_spec.SetField(task.FieldPushedAt, field.TypeTime, value)
	}
	if tu.mutation.PushedAtCleared() {
		_spec.ClearField(task.FieldPushedAt, field.TypeTime)
	}
	if value, ok := tu.mutation.DoneAt(); ok {
		_spec.SetField(task.FieldDoneAt, field.TypeTime, value)
	}
	if tu.mutation.DoneAtCleared() {
		_spec.ClearField(task.FieldDoneAt, field.TypeTime)
	}
	if value, ok := tu.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.Cap(); ok {
		_spec.SetField(task.FieldCap, field.TypeEnum, value)
	}
	if value, ok := tu.mutation.OutputBig(); ok {
		_spec.SetField(task.FieldOutputBig, field.TypeBool, value)
	}
	if tu.mutation.OutputBigCleared() {
		_spec.ClearField(task.FieldOutputBig, field.TypeBool)
	}
	if tu.mutation.CommandCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CommandIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.AntCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.AntTable,
			Columns: []string{task.AntColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ant.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.AntIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.AntTable,
			Columns: []string{task.AntColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ant.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.BlobberArgsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.BlobberArgsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.BlobberOutputCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.BlobberOutputIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TaskMutation
}

// SetCommandID sets the "command_id" field.
func (tuo *TaskUpdateOne) SetCommandID(i int64) *TaskUpdateOne {
	tuo.mutation.SetCommandID(i)
	return tuo
}

// SetNillableCommandID sets the "command_id" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCommandID(i *int64) *TaskUpdateOne {
	if i != nil {
		tuo.SetCommandID(*i)
	}
	return tuo
}

// SetAntID sets the "ant_id" field.
func (tuo *TaskUpdateOne) SetAntID(u uint32) *TaskUpdateOne {
	tuo.mutation.SetAntID(u)
	return tuo
}

// SetNillableAntID sets the "ant_id" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableAntID(u *uint32) *TaskUpdateOne {
	if u != nil {
		tuo.SetAntID(*u)
	}
	return tuo
}

// SetCreatedAt sets the "created_at" field.
func (tuo *TaskUpdateOne) SetCreatedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCreatedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetPushedAt sets the "pushed_at" field.
func (tuo *TaskUpdateOne) SetPushedAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetPushedAt(t)
	return tuo
}

// SetNillablePushedAt sets the "pushed_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillablePushedAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetPushedAt(*t)
	}
	return tuo
}

// ClearPushedAt clears the value of the "pushed_at" field.
func (tuo *TaskUpdateOne) ClearPushedAt() *TaskUpdateOne {
	tuo.mutation.ClearPushedAt()
	return tuo
}

// SetDoneAt sets the "done_at" field.
func (tuo *TaskUpdateOne) SetDoneAt(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetDoneAt(t)
	return tuo
}

// SetNillableDoneAt sets the "done_at" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDoneAt(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetDoneAt(*t)
	}
	return tuo
}

// ClearDoneAt clears the value of the "done_at" field.
func (tuo *TaskUpdateOne) ClearDoneAt() *TaskUpdateOne {
	tuo.mutation.ClearDoneAt()
	return tuo
}

// SetStatus sets the "status" field.
func (tuo *TaskUpdateOne) SetStatus(ss shared.TaskStatus) *TaskUpdateOne {
	tuo.mutation.SetStatus(ss)
	return tuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableStatus(ss *shared.TaskStatus) *TaskUpdateOne {
	if ss != nil {
		tuo.SetStatus(*ss)
	}
	return tuo
}

// SetCap sets the "cap" field.
func (tuo *TaskUpdateOne) SetCap(s shared.Capability) *TaskUpdateOne {
	tuo.mutation.SetCap(s)
	return tuo
}

// SetNillableCap sets the "cap" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableCap(s *shared.Capability) *TaskUpdateOne {
	if s != nil {
		tuo.SetCap(*s)
	}
	return tuo
}

// SetArgsID sets the "args_id" field.
func (tuo *TaskUpdateOne) SetArgsID(i int) *TaskUpdateOne {
	tuo.mutation.SetArgsID(i)
	return tuo
}

// SetNillableArgsID sets the "args_id" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableArgsID(i *int) *TaskUpdateOne {
	if i != nil {
		tuo.SetArgsID(*i)
	}
	return tuo
}

// SetOutputID sets the "output_id" field.
func (tuo *TaskUpdateOne) SetOutputID(i int) *TaskUpdateOne {
	tuo.mutation.SetOutputID(i)
	return tuo
}

// SetNillableOutputID sets the "output_id" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOutputID(i *int) *TaskUpdateOne {
	if i != nil {
		tuo.SetOutputID(*i)
	}
	return tuo
}

// ClearOutputID clears the value of the "output_id" field.
func (tuo *TaskUpdateOne) ClearOutputID() *TaskUpdateOne {
	tuo.mutation.ClearOutputID()
	return tuo
}

// SetOutputBig sets the "output_big" field.
func (tuo *TaskUpdateOne) SetOutputBig(b bool) *TaskUpdateOne {
	tuo.mutation.SetOutputBig(b)
	return tuo
}

// SetNillableOutputBig sets the "output_big" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOutputBig(b *bool) *TaskUpdateOne {
	if b != nil {
		tuo.SetOutputBig(*b)
	}
	return tuo
}

// ClearOutputBig clears the value of the "output_big" field.
func (tuo *TaskUpdateOne) ClearOutputBig() *TaskUpdateOne {
	tuo.mutation.ClearOutputBig()
	return tuo
}

// SetCommand sets the "command" edge to the Command entity.
func (tuo *TaskUpdateOne) SetCommand(c *Command) *TaskUpdateOne {
	return tuo.SetCommandID(c.ID)
}

// SetAnt sets the "ant" edge to the Ant entity.
func (tuo *TaskUpdateOne) SetAnt(a *Ant) *TaskUpdateOne {
	return tuo.SetAntID(a.ID)
}

// SetBlobberArgsID sets the "blobber_args" edge to the Blobber entity by ID.
func (tuo *TaskUpdateOne) SetBlobberArgsID(id int) *TaskUpdateOne {
	tuo.mutation.SetBlobberArgsID(id)
	return tuo
}

// SetBlobberArgs sets the "blobber_args" edge to the Blobber entity.
func (tuo *TaskUpdateOne) SetBlobberArgs(b *Blobber) *TaskUpdateOne {
	return tuo.SetBlobberArgsID(b.ID)
}

// SetBlobberOutputID sets the "blobber_output" edge to the Blobber entity by ID.
func (tuo *TaskUpdateOne) SetBlobberOutputID(id int) *TaskUpdateOne {
	tuo.mutation.SetBlobberOutputID(id)
	return tuo
}

// SetNillableBlobberOutputID sets the "blobber_output" edge to the Blobber entity by ID if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableBlobberOutputID(id *int) *TaskUpdateOne {
	if id != nil {
		tuo = tuo.SetBlobberOutputID(*id)
	}
	return tuo
}

// SetBlobberOutput sets the "blobber_output" edge to the Blobber entity.
func (tuo *TaskUpdateOne) SetBlobberOutput(b *Blobber) *TaskUpdateOne {
	return tuo.SetBlobberOutputID(b.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// ClearCommand clears the "command" edge to the Command entity.
func (tuo *TaskUpdateOne) ClearCommand() *TaskUpdateOne {
	tuo.mutation.ClearCommand()
	return tuo
}

// ClearAnt clears the "ant" edge to the Ant entity.
func (tuo *TaskUpdateOne) ClearAnt() *TaskUpdateOne {
	tuo.mutation.ClearAnt()
	return tuo
}

// ClearBlobberArgs clears the "blobber_args" edge to the Blobber entity.
func (tuo *TaskUpdateOne) ClearBlobberArgs() *TaskUpdateOne {
	tuo.mutation.ClearBlobberArgs()
	return tuo
}

// ClearBlobberOutput clears the "blobber_output" edge to the Blobber entity.
func (tuo *TaskUpdateOne) ClearBlobberOutput() *TaskUpdateOne {
	tuo.mutation.ClearBlobberOutput()
	return tuo
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Task.status": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Cap(); ok {
		if err := task.CapValidator(v); err != nil {
			return &ValidationError{Name: "cap", err: fmt.Errorf(`ent: validator failed for field "Task.cap": %w`, err)}
		}
	}
	if tuo.mutation.CommandCleared() && len(tuo.mutation.CommandIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Task.command"`)
	}
	if tuo.mutation.AntCleared() && len(tuo.mutation.AntIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Task.ant"`)
	}
	if tuo.mutation.BlobberArgsCleared() && len(tuo.mutation.BlobberArgsIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Task.blobber_args"`)
	}
	return nil
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(task.Table, task.Columns, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt64))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, task.FieldID)
		for _, f := range fields {
			if !task.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != task.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(task.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.PushedAt(); ok {
		_spec.SetField(task.FieldPushedAt, field.TypeTime, value)
	}
	if tuo.mutation.PushedAtCleared() {
		_spec.ClearField(task.FieldPushedAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.DoneAt(); ok {
		_spec.SetField(task.FieldDoneAt, field.TypeTime, value)
	}
	if tuo.mutation.DoneAtCleared() {
		_spec.ClearField(task.FieldDoneAt, field.TypeTime)
	}
	if value, ok := tuo.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.Cap(); ok {
		_spec.SetField(task.FieldCap, field.TypeEnum, value)
	}
	if value, ok := tuo.mutation.OutputBig(); ok {
		_spec.SetField(task.FieldOutputBig, field.TypeBool, value)
	}
	if tuo.mutation.OutputBigCleared() {
		_spec.ClearField(task.FieldOutputBig, field.TypeBool)
	}
	if tuo.mutation.CommandCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CommandIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.AntCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.AntTable,
			Columns: []string{task.AntColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ant.FieldID, field.TypeUint32),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.AntIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.AntTable,
			Columns: []string{task.AntColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ant.FieldID, field.TypeUint32),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.BlobberArgsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.BlobberArgsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.BlobberOutputCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.BlobberOutputIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}