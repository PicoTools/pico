// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/PicoTools/pico/internal/ent/chat"
	"github.com/PicoTools/pico/internal/ent/command"
	"github.com/PicoTools/pico/internal/ent/operator"
	"github.com/PicoTools/pico/internal/ent/predicate"
)

// OperatorQuery is the builder for querying Operator entities.
type OperatorQuery struct {
	config
	ctx         *QueryContext
	order       []operator.OrderOption
	inters      []Interceptor
	predicates  []predicate.Operator
	withChat    *ChatQuery
	withCommand *CommandQuery
	modifiers   []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OperatorQuery builder.
func (oq *OperatorQuery) Where(ps ...predicate.Operator) *OperatorQuery {
	oq.predicates = append(oq.predicates, ps...)
	return oq
}

// Limit the number of records to be returned by this query.
func (oq *OperatorQuery) Limit(limit int) *OperatorQuery {
	oq.ctx.Limit = &limit
	return oq
}

// Offset to start from.
func (oq *OperatorQuery) Offset(offset int) *OperatorQuery {
	oq.ctx.Offset = &offset
	return oq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oq *OperatorQuery) Unique(unique bool) *OperatorQuery {
	oq.ctx.Unique = &unique
	return oq
}

// Order specifies how the records should be ordered.
func (oq *OperatorQuery) Order(o ...operator.OrderOption) *OperatorQuery {
	oq.order = append(oq.order, o...)
	return oq
}

// QueryChat chains the current query on the "chat" edge.
func (oq *OperatorQuery) QueryChat() *ChatQuery {
	query := (&ChatClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(operator.Table, operator.FieldID, selector),
			sqlgraph.To(chat.Table, chat.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, operator.ChatTable, operator.ChatColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryCommand chains the current query on the "command" edge.
func (oq *OperatorQuery) QueryCommand() *CommandQuery {
	query := (&CommandClient{config: oq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(operator.Table, operator.FieldID, selector),
			sqlgraph.To(command.Table, command.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, operator.CommandTable, operator.CommandColumn),
		)
		fromU = sqlgraph.SetNeighbors(oq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Operator entity from the query.
// Returns a *NotFoundError when no Operator was found.
func (oq *OperatorQuery) First(ctx context.Context) (*Operator, error) {
	nodes, err := oq.Limit(1).All(setContextOp(ctx, oq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{operator.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oq *OperatorQuery) FirstX(ctx context.Context) *Operator {
	node, err := oq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Operator ID from the query.
// Returns a *NotFoundError when no Operator ID was found.
func (oq *OperatorQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = oq.Limit(1).IDs(setContextOp(ctx, oq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{operator.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oq *OperatorQuery) FirstIDX(ctx context.Context) int64 {
	id, err := oq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Operator entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Operator entity is found.
// Returns a *NotFoundError when no Operator entities are found.
func (oq *OperatorQuery) Only(ctx context.Context) (*Operator, error) {
	nodes, err := oq.Limit(2).All(setContextOp(ctx, oq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{operator.Label}
	default:
		return nil, &NotSingularError{operator.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oq *OperatorQuery) OnlyX(ctx context.Context) *Operator {
	node, err := oq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Operator ID in the query.
// Returns a *NotSingularError when more than one Operator ID is found.
// Returns a *NotFoundError when no entities are found.
func (oq *OperatorQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = oq.Limit(2).IDs(setContextOp(ctx, oq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{operator.Label}
	default:
		err = &NotSingularError{operator.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oq *OperatorQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := oq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Operators.
func (oq *OperatorQuery) All(ctx context.Context) ([]*Operator, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryAll)
	if err := oq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Operator, *OperatorQuery]()
	return withInterceptors[[]*Operator](ctx, oq, qr, oq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oq *OperatorQuery) AllX(ctx context.Context) []*Operator {
	nodes, err := oq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Operator IDs.
func (oq *OperatorQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if oq.ctx.Unique == nil && oq.path != nil {
		oq.Unique(true)
	}
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryIDs)
	if err = oq.Select(operator.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oq *OperatorQuery) IDsX(ctx context.Context) []int64 {
	ids, err := oq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oq *OperatorQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryCount)
	if err := oq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oq, querierCount[*OperatorQuery](), oq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oq *OperatorQuery) CountX(ctx context.Context) int {
	count, err := oq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oq *OperatorQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oq.ctx, ent.OpQueryExist)
	switch _, err := oq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oq *OperatorQuery) ExistX(ctx context.Context) bool {
	exist, err := oq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OperatorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oq *OperatorQuery) Clone() *OperatorQuery {
	if oq == nil {
		return nil
	}
	return &OperatorQuery{
		config:      oq.config,
		ctx:         oq.ctx.Clone(),
		order:       append([]operator.OrderOption{}, oq.order...),
		inters:      append([]Interceptor{}, oq.inters...),
		predicates:  append([]predicate.Operator{}, oq.predicates...),
		withChat:    oq.withChat.Clone(),
		withCommand: oq.withCommand.Clone(),
		// clone intermediate query.
		sql:  oq.sql.Clone(),
		path: oq.path,
	}
}

// WithChat tells the query-builder to eager-load the nodes that are connected to
// the "chat" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OperatorQuery) WithChat(opts ...func(*ChatQuery)) *OperatorQuery {
	query := (&ChatClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withChat = query
	return oq
}

// WithCommand tells the query-builder to eager-load the nodes that are connected to
// the "command" edge. The optional arguments are used to configure the query builder of the edge.
func (oq *OperatorQuery) WithCommand(opts ...func(*CommandQuery)) *OperatorQuery {
	query := (&CommandClient{config: oq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	oq.withCommand = query
	return oq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Operator.Query().
//		GroupBy(operator.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oq *OperatorQuery) GroupBy(field string, fields ...string) *OperatorGroupBy {
	oq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OperatorGroupBy{build: oq}
	grbuild.flds = &oq.ctx.Fields
	grbuild.label = operator.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Operator.Query().
//		Select(operator.FieldCreatedAt).
//		Scan(ctx, &v)
func (oq *OperatorQuery) Select(fields ...string) *OperatorSelect {
	oq.ctx.Fields = append(oq.ctx.Fields, fields...)
	sbuild := &OperatorSelect{OperatorQuery: oq}
	sbuild.label = operator.Label
	sbuild.flds, sbuild.scan = &oq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OperatorSelect configured with the given aggregations.
func (oq *OperatorQuery) Aggregate(fns ...AggregateFunc) *OperatorSelect {
	return oq.Select().Aggregate(fns...)
}

func (oq *OperatorQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oq); err != nil {
				return err
			}
		}
	}
	for _, f := range oq.ctx.Fields {
		if !operator.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oq.path != nil {
		prev, err := oq.path(ctx)
		if err != nil {
			return err
		}
		oq.sql = prev
	}
	return nil
}

func (oq *OperatorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Operator, error) {
	var (
		nodes       = []*Operator{}
		_spec       = oq.querySpec()
		loadedTypes = [2]bool{
			oq.withChat != nil,
			oq.withCommand != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Operator).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Operator{config: oq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(oq.modifiers) > 0 {
		_spec.Modifiers = oq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oq.withChat; query != nil {
		if err := oq.loadChat(ctx, query, nodes,
			func(n *Operator) { n.Edges.Chat = []*Chat{} },
			func(n *Operator, e *Chat) { n.Edges.Chat = append(n.Edges.Chat, e) }); err != nil {
			return nil, err
		}
	}
	if query := oq.withCommand; query != nil {
		if err := oq.loadCommand(ctx, query, nodes,
			func(n *Operator) { n.Edges.Command = []*Command{} },
			func(n *Operator, e *Command) { n.Edges.Command = append(n.Edges.Command, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oq *OperatorQuery) loadChat(ctx context.Context, query *ChatQuery, nodes []*Operator, init func(*Operator), assign func(*Operator, *Chat)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Operator)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(chat.FieldAuthorID)
	}
	query.Where(predicate.Chat(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(operator.ChatColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AuthorID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "author_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (oq *OperatorQuery) loadCommand(ctx context.Context, query *CommandQuery, nodes []*Operator, init func(*Operator), assign func(*Operator, *Command)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int64]*Operator)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(command.FieldAuthorID)
	}
	query.Where(predicate.Command(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(operator.CommandColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AuthorID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "author_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (oq *OperatorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oq.querySpec()
	if len(oq.modifiers) > 0 {
		_spec.Modifiers = oq.modifiers
	}
	_spec.Node.Columns = oq.ctx.Fields
	if len(oq.ctx.Fields) > 0 {
		_spec.Unique = oq.ctx.Unique != nil && *oq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oq.driver, _spec)
}

func (oq *OperatorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(operator.Table, operator.Columns, sqlgraph.NewFieldSpec(operator.FieldID, field.TypeInt64))
	_spec.From = oq.sql
	if unique := oq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oq.path != nil {
		_spec.Unique = true
	}
	if fields := oq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, operator.FieldID)
		for i := range fields {
			if fields[i] != operator.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oq *OperatorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oq.driver.Dialect())
	t1 := builder.Table(operator.Table)
	columns := oq.ctx.Fields
	if len(columns) == 0 {
		columns = operator.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oq.sql != nil {
		selector = oq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oq.ctx.Unique != nil && *oq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range oq.modifiers {
		m(selector)
	}
	for _, p := range oq.predicates {
		p(selector)
	}
	for _, p := range oq.order {
		p(selector)
	}
	if offset := oq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (oq *OperatorQuery) ForUpdate(opts ...sql.LockOption) *OperatorQuery {
	if oq.driver.Dialect() == dialect.Postgres {
		oq.Unique(false)
	}
	oq.modifiers = append(oq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return oq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (oq *OperatorQuery) ForShare(opts ...sql.LockOption) *OperatorQuery {
	if oq.driver.Dialect() == dialect.Postgres {
		oq.Unique(false)
	}
	oq.modifiers = append(oq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return oq
}

// OperatorGroupBy is the group-by builder for Operator entities.
type OperatorGroupBy struct {
	selector
	build *OperatorQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ogb *OperatorGroupBy) Aggregate(fns ...AggregateFunc) *OperatorGroupBy {
	ogb.fns = append(ogb.fns, fns...)
	return ogb
}

// Scan applies the selector query and scans the result into the given value.
func (ogb *OperatorGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ogb.build.ctx, ent.OpQueryGroupBy)
	if err := ogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OperatorQuery, *OperatorGroupBy](ctx, ogb.build, ogb, ogb.build.inters, v)
}

func (ogb *OperatorGroupBy) sqlScan(ctx context.Context, root *OperatorQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ogb.fns))
	for _, fn := range ogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ogb.flds)+len(ogb.fns))
		for _, f := range *ogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OperatorSelect is the builder for selecting fields of Operator entities.
type OperatorSelect struct {
	*OperatorQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (os *OperatorSelect) Aggregate(fns ...AggregateFunc) *OperatorSelect {
	os.fns = append(os.fns, fns...)
	return os
}

// Scan applies the selector query and scans the result into the given value.
func (os *OperatorSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, os.ctx, ent.OpQuerySelect)
	if err := os.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OperatorQuery, *OperatorSelect](ctx, os.OperatorQuery, os, os.inters, v)
}

func (os *OperatorSelect) sqlScan(ctx context.Context, root *OperatorQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(os.fns))
	for _, fn := range os.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*os.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := os.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
