package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"math/rand"
	"time"

	"github.com/ryansheehan/tagger/graph/generated"
	"github.com/ryansheehan/tagger/graph/model"
)

func (r *mutationResolver) CreateMessage(ctx context.Context, input model.NewTagMsg) (*model.TagMsg, error) {
	var postAt time.Time
	var currentTime = time.Now().UTC()
	if input.PostAt == nil {
		postAt = currentTime
	}
	msg := &model.TagMsg{
		ID:        rand.Int63(),
		Text:      input.Text,
		Geo:       input.Geo,
		PostAt:    postAt,
		CreatedAt: currentTime,
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
