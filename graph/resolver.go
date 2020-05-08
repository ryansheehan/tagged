//go:generate go run github.com/99designs/gqlgen

package graph

import "github.com/ryansheehan/tagger/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver resolves the root ql objects
type Resolver struct {
	messages []*model.TagMsg
}
