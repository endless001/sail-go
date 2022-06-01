package core

import "context"

type Route struct {
	ID int64 `db:"id" json:"id"`
}

type RouteStore interface {
	// Find returns a build from the datastore.
	Find(context.Context, int64) (*Route, error)
}
