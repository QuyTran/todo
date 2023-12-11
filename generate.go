package todo

//go:generate go run -mod=mod ./repositories/ent/ent/entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen -config=api/graphql/gqlgen.yml
