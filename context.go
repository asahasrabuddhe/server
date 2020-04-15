package server

import "context"

// Context is a type passed to each handler
type Context struct {
	context.Context
}

