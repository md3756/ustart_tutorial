package storage

import (
	"context"

	// "github.com/md3756/ustart_tutorial/record/recordpb"
)

// Storage is a database-agnostic interface for persisting customer data
type Storage interface {
	// New(context.Context, string, string, string, string) error
	History(context.Context, string) (recordpb.Record, error)
	Pay(context.Context, []string, bool, map[string][]string, string) ([]string, error)
	Insert(ctx context.Context, string, string, string, int) error

	ErrRecordDoesNotExist() error
	ErrTooManyResults() error
}
