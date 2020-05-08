package model

import (
	"time"

	"github.com/ryansheehan/tagger/graph/scalar"
)

// TagMsg is the graphql implementation of the model of the same name
type TagMsg struct {
	ID        int64                  `json:"id"`
	Text      string                 `json:"text"`
	PostAt    time.Time              `json:"postAt"`
	CreatedAt time.Time              `json:"createdAt"`
	Geo       *scalar.GeoLoc         `json:"geo"`
	Tags      map[string]interface{} `json:"tags"`
}

// NewTagMsg is the graphql implementation for input of a new TagMsg
type NewTagMsg struct {
	Text   string         `json:"text"`
	PostAt *time.Time     `json:"postAt"`
	Geo    *scalar.GeoLoc `json:"geo"`
}
