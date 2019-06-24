package storage

import (
	"context"
)

// Storage is a database-agnostic interface for persisting customer data
type Storage interface {
	Register(context.Context, string, string, string, string, string, string) error
	Search(context.Context, []string, bool, map[string][]string, string) ([]string, error)
	Lookup(context.Context, string) (bool, error)
	UpdateAvailable(context.Context, string, bool) error
	CheckAvailable(context.Context, string) (bool, error)
	// rest of the functions
	ErrCarDoesNotExist() error
	ErrTooManyResults() error
}
