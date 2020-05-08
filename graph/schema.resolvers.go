package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/ryansheehan/tagger/graph/generated"
	"github.com/ryansheehan/tagger/graph/model"
	"github.com/ryansheehan/tagger/parser"
)

func (r *mutationResolver) CreateMessage(ctx context.Context, input model.NewTagMsg) (*model.TagMsg, error) {
	trimmedText := strings.TrimSpace(input.Text)
	if trimmedText == "" {
		return nil, fmt.Errorf("text is required, and must be more than whitespace")
	}
	
	var postAt time.Time
	var currentTime = time.Now().UTC()
	if input.PostAt == nil {
		postAt = currentTime
	}
	tags := parser.Parse(trimmedText)
	msg := &model.TagMsg{
		ID:        rand.Int63(),
		Text:      trimmedText,
		Geo:       input.Geo,
		PostAt:    postAt,
		CreatedAt: currentTime,
		Tags:      tags,
	}

	r.messages = append(r.messages, msg)
	return msg, nil
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.TagMsg, error) {
	return r.messages, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
