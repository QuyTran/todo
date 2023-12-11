// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"todo/repositories/ent/ent/predicate"
	"todo/repositories/ent/ent/todo"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTodo = "Todo"
)

// TodoMutation represents an operation that mutates the Todo nodes in the graph.
type TodoMutation struct {
	config
	op              Op
	typ             string
	id              *int
	text            *string
	status          *todo.Status
	priority        *int
	addpriority     *int
	clearedFields   map[string]struct{}
	children        map[int]struct{}
	removedchildren map[int]struct{}
	clearedchildren bool
	parent          *int
	clearedparent   bool
	done            bool
	oldValue        func(context.Context) (*Todo, error)
	predicates      []predicate.Todo
}

var _ ent.Mutation = (*TodoMutation)(nil)

// todoOption allows management of the mutation configuration using functional options.
type todoOption func(*TodoMutation)

// newTodoMutation creates new mutation for the Todo entity.
func newTodoMutation(c config, op Op, opts ...todoOption) *TodoMutation {
	m := &TodoMutation{
		config:        c,
		op:            op,
		typ:           TypeTodo,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTodoID sets the ID field of the mutation.
func withTodoID(id int) todoOption {
	return func(m *TodoMutation) {
		var (
			err   error
			once  sync.Once
			value *Todo
		)
		m.oldValue = func(ctx context.Context) (*Todo, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Todo.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTodo sets the old Todo of the mutation.
func withTodo(node *Todo) todoOption {
	return func(m *TodoMutation) {
		m.oldValue = func(context.Context) (*Todo, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TodoMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TodoMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TodoMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TodoMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Todo.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetText sets the "text" field.
func (m *TodoMutation) SetText(s string) {
	m.text = &s
}

// Text returns the value of the "text" field in the mutation.
func (m *TodoMutation) Text() (r string, exists bool) {
	v := m.text
	if v == nil {
		return
	}
	return *v, true
}

// OldText returns the old "text" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldText(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldText is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldText requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldText: %w", err)
	}
	return oldValue.Text, nil
}

// ResetText resets all changes to the "text" field.
func (m *TodoMutation) ResetText() {
	m.text = nil
}

// SetStatus sets the "status" field.
func (m *TodoMutation) SetStatus(t todo.Status) {
	m.status = &t
}

// Status returns the value of the "status" field in the mutation.
func (m *TodoMutation) Status() (r todo.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldStatus(ctx context.Context) (v todo.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *TodoMutation) ResetStatus() {
	m.status = nil
}

// SetPriority sets the "priority" field.
func (m *TodoMutation) SetPriority(i int) {
	m.priority = &i
	m.addpriority = nil
}

// Priority returns the value of the "priority" field in the mutation.
func (m *TodoMutation) Priority() (r int, exists bool) {
	v := m.priority
	if v == nil {
		return
	}
	return *v, true
}

// OldPriority returns the old "priority" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldPriority(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPriority is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPriority requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPriority: %w", err)
	}
	return oldValue.Priority, nil
}

// AddPriority adds i to the "priority" field.
func (m *TodoMutation) AddPriority(i int) {
	if m.addpriority != nil {
		*m.addpriority += i
	} else {
		m.addpriority = &i
	}
}

// AddedPriority returns the value that was added to the "priority" field in this mutation.
func (m *TodoMutation) AddedPriority() (r int, exists bool) {
	v := m.addpriority
	if v == nil {
		return
	}
	return *v, true
}

// ResetPriority resets all changes to the "priority" field.
func (m *TodoMutation) ResetPriority() {
	m.priority = nil
	m.addpriority = nil
}

// AddChildIDs adds the "children" edge to the Todo entity by ids.
func (m *TodoMutation) AddChildIDs(ids ...int) {
	if m.children == nil {
		m.children = make(map[int]struct{})
	}
	for i := range ids {
		m.children[ids[i]] = struct{}{}
	}
}

// ClearChildren clears the "children" edge to the Todo entity.
func (m *TodoMutation) ClearChildren() {
	m.clearedchildren = true
}

// ChildrenCleared reports if the "children" edge to the Todo entity was cleared.
func (m *TodoMutation) ChildrenCleared() bool {
	return m.clearedchildren
}

// RemoveChildIDs removes the "children" edge to the Todo entity by IDs.
func (m *TodoMutation) RemoveChildIDs(ids ...int) {
	if m.removedchildren == nil {
		m.removedchildren = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.children, ids[i])
		m.removedchildren[ids[i]] = struct{}{}
	}
}

// RemovedChildren returns the removed IDs of the "children" edge to the Todo entity.
func (m *TodoMutation) RemovedChildrenIDs() (ids []int) {
	for id := range m.removedchildren {
		ids = append(ids, id)
	}
	return
}

// ChildrenIDs returns the "children" edge IDs in the mutation.
func (m *TodoMutation) ChildrenIDs() (ids []int) {
	for id := range m.children {
		ids = append(ids, id)
	}
	return
}

// ResetChildren resets all changes to the "children" edge.
func (m *TodoMutation) ResetChildren() {
	m.children = nil
	m.clearedchildren = false
	m.removedchildren = nil
}

// SetParentID sets the "parent" edge to the Todo entity by id.
func (m *TodoMutation) SetParentID(id int) {
	m.parent = &id
}

// ClearParent clears the "parent" edge to the Todo entity.
func (m *TodoMutation) ClearParent() {
	m.clearedparent = true
}

// ParentCleared reports if the "parent" edge to the Todo entity was cleared.
func (m *TodoMutation) ParentCleared() bool {
	return m.clearedparent
}

// ParentID returns the "parent" edge ID in the mutation.
func (m *TodoMutation) ParentID() (id int, exists bool) {
	if m.parent != nil {
		return *m.parent, true
	}
	return
}

// ParentIDs returns the "parent" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// ParentID instead. It exists only for internal usage by the builders.
func (m *TodoMutation) ParentIDs() (ids []int) {
	if id := m.parent; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetParent resets all changes to the "parent" edge.
func (m *TodoMutation) ResetParent() {
	m.parent = nil
	m.clearedparent = false
}

// Where appends a list predicates to the TodoMutation builder.
func (m *TodoMutation) Where(ps ...predicate.Todo) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TodoMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TodoMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Todo, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TodoMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TodoMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Todo).
func (m *TodoMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TodoMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.text != nil {
		fields = append(fields, todo.FieldText)
	}
	if m.status != nil {
		fields = append(fields, todo.FieldStatus)
	}
	if m.priority != nil {
		fields = append(fields, todo.FieldPriority)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TodoMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case todo.FieldText:
		return m.Text()
	case todo.FieldStatus:
		return m.Status()
	case todo.FieldPriority:
		return m.Priority()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TodoMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case todo.FieldText:
		return m.OldText(ctx)
	case todo.FieldStatus:
		return m.OldStatus(ctx)
	case todo.FieldPriority:
		return m.OldPriority(ctx)
	}
	return nil, fmt.Errorf("unknown Todo field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TodoMutation) SetField(name string, value ent.Value) error {
	switch name {
	case todo.FieldText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetText(v)
		return nil
	case todo.FieldStatus:
		v, ok := value.(todo.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case todo.FieldPriority:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPriority(v)
		return nil
	}
	return fmt.Errorf("unknown Todo field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TodoMutation) AddedFields() []string {
	var fields []string
	if m.addpriority != nil {
		fields = append(fields, todo.FieldPriority)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TodoMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case todo.FieldPriority:
		return m.AddedPriority()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TodoMutation) AddField(name string, value ent.Value) error {
	switch name {
	case todo.FieldPriority:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPriority(v)
		return nil
	}
	return fmt.Errorf("unknown Todo numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TodoMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TodoMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TodoMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Todo nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TodoMutation) ResetField(name string) error {
	switch name {
	case todo.FieldText:
		m.ResetText()
		return nil
	case todo.FieldStatus:
		m.ResetStatus()
		return nil
	case todo.FieldPriority:
		m.ResetPriority()
		return nil
	}
	return fmt.Errorf("unknown Todo field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TodoMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.children != nil {
		edges = append(edges, todo.EdgeChildren)
	}
	if m.parent != nil {
		edges = append(edges, todo.EdgeParent)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TodoMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case todo.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.children))
		for id := range m.children {
			ids = append(ids, id)
		}
		return ids
	case todo.EdgeParent:
		if id := m.parent; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TodoMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	if m.removedchildren != nil {
		edges = append(edges, todo.EdgeChildren)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TodoMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case todo.EdgeChildren:
		ids := make([]ent.Value, 0, len(m.removedchildren))
		for id := range m.removedchildren {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TodoMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.clearedchildren {
		edges = append(edges, todo.EdgeChildren)
	}
	if m.clearedparent {
		edges = append(edges, todo.EdgeParent)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TodoMutation) EdgeCleared(name string) bool {
	switch name {
	case todo.EdgeChildren:
		return m.clearedchildren
	case todo.EdgeParent:
		return m.clearedparent
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TodoMutation) ClearEdge(name string) error {
	switch name {
	case todo.EdgeParent:
		m.ClearParent()
		return nil
	}
	return fmt.Errorf("unknown Todo unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TodoMutation) ResetEdge(name string) error {
	switch name {
	case todo.EdgeChildren:
		m.ResetChildren()
		return nil
	case todo.EdgeParent:
		m.ResetParent()
		return nil
	}
	return fmt.Errorf("unknown Todo edge %s", name)
}
