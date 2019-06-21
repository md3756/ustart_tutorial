package record

import (
	"context"

	"github.com/sea350/ustart_tutorial/record/recordpb"
)

// History ...
func (record *Record) History(ctx context.Context, req *recordpb.HistoryRequest) (*recordpb.HistoryResponse, error) {

	rec, err := record.strg.Lookup(ctx, req.CarID)
	if err != nil {
		return nil, err
	}

	return &recordpb.HistoryResponse{RecordProfile: &rec}, nil

}
