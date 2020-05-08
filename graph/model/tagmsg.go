package model

import (
	"github.com/ryansheehan/tagger/graph/scalar"
)

// TagMsg is the graphql implementation of the model of the same name
type TagMsg struct {
	ID   int64         `json:"id"`
	Text string        `json:"text"`
	Geo  *scalar.GeoLoc `json:"geo"`
}

// NewTagMsg is the graphql implementation for input of a new TagMsg
type NewTagMsg struct {
	Text string        `json:"text"`
	Geo  *scalar.GeoLoc `json:"geo"`
}
