package storage

import (
	"context"

	"github.com/md3756/ustart_tutorial/record/recordpb"
)

// Storage is a database-agnostic interface for persisting customer data
type Storage interface {
	Register(context.Context, string, string) error
	Lookup(context.Context, string) (recordpb.Record, error)
	Search(context.Context, []string, bool, map[string][]string, string) ([]string, error)
	// rest of the functions
	ErrRecordDoesNotExist() error
	ErrTooManyResults() error
}
