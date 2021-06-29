package listing

import (
	"time"
)

// Institution defines the properties of a Institutions to be listed
type Institution struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created"`
}
