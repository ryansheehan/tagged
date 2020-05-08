package scalar

import (
	"encoding/json"
	"fmt"
	"io"
)

// GeoLoc stores latitudinal and longitudel geospacial coordinates
type GeoLoc struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// UnmarshalGQL implements the graphql.Unmarshaler interface
func (geo *GeoLoc) UnmarshalGQL(v interface{}) error {
	if m, ok := v.(map[string]interface{}); ok {
		lat, latOk := m["lat"]
		lng, lngOk := m["lng"]
		if latOk && lngOk {
			geo.Lat = lat.(float64)
			geo.Lng = lng.(float64)
			return nil
		}
	}
	return fmt.Errorf("%T is not a GeoLocation", v)
}

// MarshalGQL implements the graphql.Marshaler interface
func (geo GeoLoc) MarshalGQL(w io.Writer) {
	err := json.NewEncoder(w).Encode(geo)
	if err != nil {
		panic(err)
	}
}
