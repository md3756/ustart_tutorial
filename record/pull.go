package record

import (
	"context"
	// "github.com/md3756/ustart_tutorial/record/recordpb"
)

// Pull retreives all customer data based off of a username
func (record *Record) Pull(ctx context.Context, req *recordpb.PullRequest) (*recordpb.PullResponse, error) {

	rec, err := record.strg.Lookup(ctx, req.CarID)
	if err != nil {
		return nil, err
	}

	return &recordpb.PullResponse{RecordProfile: &rec}, nil

}
