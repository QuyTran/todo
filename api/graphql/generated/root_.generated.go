// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"bytes"
	"context"
	"errors"
	"sync/atomic"
	"todo/repositories/ent/ent"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Mutation struct {
		CreateTodo func(childComplexity int, input ent.CreateTodoInput) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Query struct {
		Node  func(childComplexity int, id int) int
		Nodes func(childComplexity int, ids []int) int
		Todos func(childComplexity int, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.TodoOrder, where *ent.TodoWhereInput) int
	}

	Todo struct {
		Children func(childComplexity int) int
		ID       func(childComplexity int) int
		Parent   func(childComplexity int) int
		Priority func(childComplexity int) int
		Status   func(childComplexity int) int
		Text     func(childComplexity int) int
	}

	TodoConnection struct {
		Edges      func(childComplexity int) int
		PageInfo   func(childComplexity int) int
		TotalCount func(childComplexity int) int
	}

	TodoEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "Mutation.createTodo":
		if e.complexity.Mutation.CreateTodo == nil {
			break
		}

		args, err := ec.field_Mutation_createTodo_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateTodo(childComplexity, args["input"].(ent.CreateTodoInput)), true

	case "PageInfo.endCursor":
		if e.complexity.PageInfo.EndCursor == nil {
			break
		}

		return e.complexity.PageInfo.EndCursor(childComplexity), true

	case "PageInfo.hasNextPage":
		if e.complexity.PageInfo.HasNextPage == nil {
			break
		}

		return e.complexity.PageInfo.HasNextPage(childComplexity), true

	case "PageInfo.hasPreviousPage":
		if e.complexity.PageInfo.HasPreviousPage == nil {
			break
		}

		return e.complexity.PageInfo.HasPreviousPage(childComplexity), true

	case "PageInfo.startCursor":
		if e.complexity.PageInfo.StartCursor == nil {
			break
		}

		return e.complexity.PageInfo.StartCursor(childComplexity), true

	case "Query.node":
		if e.complexity.Query.Node == nil {
			break
		}

		args, err := ec.field_Query_node_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Node(childComplexity, args["id"].(int)), true

	case "Query.nodes":
		if e.complexity.Query.Nodes == nil {
			break
		}

		args, err := ec.field_Query_nodes_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Nodes(childComplexity, args["ids"].([]int)), true

	case "Query.todos":
		if e.complexity.Query.Todos == nil {
			break
		}

		args, err := ec.field_Query_todos_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Todos(childComplexity, args["after"].(*entgql.Cursor[int]), args["first"].(*int), args["before"].(*entgql.Cursor[int]), args["last"].(*int), args["orderBy"].(*ent.TodoOrder), args["where"].(*ent.TodoWhereInput)), true

	case "Todo.children":
		if e.complexity.Todo.Children == nil {
			break
		}

		return e.complexity.Todo.Children(childComplexity), true

	case "Todo.id":
		if e.complexity.Todo.ID == nil {
			break
		}

		return e.complexity.Todo.ID(childComplexity), true

	case "Todo.parent":
		if e.complexity.Todo.Parent == nil {
			break
		}

		return e.complexity.Todo.Parent(childComplexity), true

	case "Todo.priority":
		if e.complexity.Todo.Priority == nil {
			break
		}

		return e.complexity.Todo.Priority(childComplexity), true

	case "Todo.status":
		if e.complexity.Todo.Status == nil {
			break
		}

		return e.complexity.Todo.Status(childComplexity), true

	case "Todo.text":
		if e.complexity.Todo.Text == nil {
			break
		}

		return e.complexity.Todo.Text(childComplexity), true

	case "TodoConnection.edges":
		if e.complexity.TodoConnection.Edges == nil {
			break
		}

		return e.complexity.TodoConnection.Edges(childComplexity), true

	case "TodoConnection.pageInfo":
		if e.complexity.TodoConnection.PageInfo == nil {
			break
		}

		return e.complexity.TodoConnection.PageInfo(childComplexity), true

	case "TodoConnection.totalCount":
		if e.complexity.TodoConnection.TotalCount == nil {
			break
		}

		return e.complexity.TodoConnection.TotalCount(childComplexity), true

	case "TodoEdge.cursor":
		if e.complexity.TodoEdge.Cursor == nil {
			break
		}

		return e.complexity.TodoEdge.Cursor(childComplexity), true

	case "TodoEdge.node":
		if e.complexity.TodoEdge.Node == nil {
			break
		}

		return e.complexity.TodoEdge.Node(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputCreateTodoInput,
		ec.unmarshalInputTodoOrder,
		ec.unmarshalInputTodoWhereInput,
		ec.unmarshalInputUpdateTodoInput,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, rc.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../../../repositories/ent/ent.graphql", Input: `directive @goField(forceResolver: Boolean, name: String) on FIELD_DEFINITION | INPUT_FIELD_DEFINITION
directive @goModel(model: String, models: [String!]) on OBJECT | INPUT_OBJECT | SCALAR | ENUM | INTERFACE | UNION
"""
CreateTodoInput is used for create Todo object.
Input was generated by ent.
"""
input CreateTodoInput {
  text: String!
  status: TodoStatus
  priority: Int
  childIDs: [ID!]
  parentID: ID
}
"""
Define a Relay Cursor type:
https://relay.dev/graphql/connections.htm#sec-Cursor
"""
scalar Cursor
"""
An object with an ID.
Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
"""
interface Node @goModel(model: "todo/repositories/ent/ent.Noder") {
  """The id of the object."""
  id: ID!
}
"""Possible directions in which to order a list of items when provided an ` + "`" + `orderBy` + "`" + ` argument."""
enum OrderDirection {
  """Specifies an ascending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  ASC
  """Specifies a descending order for a given ` + "`" + `orderBy` + "`" + ` argument."""
  DESC
}
"""
Information about pagination in a connection.
https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
"""
type PageInfo {
  """When paginating forwards, are there more items?"""
  hasNextPage: Boolean!
  """When paginating backwards, are there more items?"""
  hasPreviousPage: Boolean!
  """When paginating backwards, the cursor to continue."""
  startCursor: Cursor
  """When paginating forwards, the cursor to continue."""
  endCursor: Cursor
}
type Query {
  """Fetches an object given its ID."""
  node(
    """ID of the object."""
    id: ID!
  ): Node
  """Lookup nodes by a list of IDs."""
  nodes(
    """The list of node IDs."""
    ids: [ID!]!
  ): [Node]!
  todos(
    """Returns the elements in the list that come after the specified cursor."""
    after: Cursor

    """Returns the first _n_ elements from the list."""
    first: Int

    """Returns the elements in the list that come before the specified cursor."""
    before: Cursor

    """Returns the last _n_ elements from the list."""
    last: Int

    """Ordering options for Todos returned from the connection."""
    orderBy: TodoOrder

    """Filtering options for Todos returned from the connection."""
    where: TodoWhereInput
  ): TodoConnection!
}
type Todo implements Node {
  id: ID!
  text: String!
  status: TodoStatus!
  priority: Int!
  children: [Todo!]
  parent: Todo
}
"""A connection to a list of items."""
type TodoConnection {
  """A list of edges."""
  edges: [TodoEdge]
  """Information to aid in pagination."""
  pageInfo: PageInfo!
  """Identifies the total count of items in the connection."""
  totalCount: Int!
}
"""An edge in a connection."""
type TodoEdge {
  """The item at the end of the edge."""
  node: Todo
  """A cursor for use in pagination."""
  cursor: Cursor!
}
"""Ordering options for Todo connections"""
input TodoOrder {
  """The ordering direction."""
  direction: OrderDirection! = ASC
  """The field by which to order Todos."""
  field: TodoOrderField!
}
"""Properties by which Todo connections can be ordered."""
enum TodoOrderField {
  TEXT
  STATUS
  PRIORITY
}
"""TodoStatus is enum for the field status"""
enum TodoStatus @goModel(model: "todo/repositories/ent/ent/todo.Status") {
  IN_PROGRESS
  COMPLETED
}
"""
TodoWhereInput is used for filtering Todo objects.
Input was generated by ent.
"""
input TodoWhereInput {
  not: TodoWhereInput
  and: [TodoWhereInput!]
  or: [TodoWhereInput!]
  """id field predicates"""
  id: ID
  idNEQ: ID
  idIn: [ID!]
  idNotIn: [ID!]
  idGT: ID
  idGTE: ID
  idLT: ID
  idLTE: ID
  """text field predicates"""
  text: String
  textNEQ: String
  textIn: [String!]
  textNotIn: [String!]
  textGT: String
  textGTE: String
  textLT: String
  textLTE: String
  textContains: String
  textHasPrefix: String
  textHasSuffix: String
  textEqualFold: String
  textContainsFold: String
  """status field predicates"""
  status: TodoStatus
  statusNEQ: TodoStatus
  statusIn: [TodoStatus!]
  statusNotIn: [TodoStatus!]
  """priority field predicates"""
  priority: Int
  priorityNEQ: Int
  priorityIn: [Int!]
  priorityNotIn: [Int!]
  priorityGT: Int
  priorityGTE: Int
  priorityLT: Int
  priorityLTE: Int
  """children edge predicates"""
  hasChildren: Boolean
  hasChildrenWith: [TodoWhereInput!]
  """parent edge predicates"""
  hasParent: Boolean
  hasParentWith: [TodoWhereInput!]
}
"""
UpdateTodoInput is used for update Todo object.
Input was generated by ent.
"""
input UpdateTodoInput {
  text: String
  status: TodoStatus
  priority: Int
  addChildIDs: [ID!]
  removeChildIDs: [ID!]
  clearChildren: Boolean
  parentID: ID
  clearParent: Boolean
}
`, BuiltIn: false},
	{Name: "../../../repositories/ent/todo.graphql", Input: `type Mutation {
    # The input and the output are types generated by Ent.
    createTodo(input: CreateTodoInput!): Todo
}`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
