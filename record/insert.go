package record

import (
	"context"

	// "github.com/md3756/ustart_tutorial/record/recordpb"
)

// Insert ...
func (record *Record) Insert(ctx context.Context, req *recordpb.InsertRequest) (*recordpb.InsertResponse, error) {
	
	intVal, err := strconv.Atoi(req.Rate)
	if err != nil {
		nil, err
	}
	err := record.strg.Insert(ctx, req.CarID, req.UserID, req.DateStart, intVal)
	if err = nil {
		return nil, err
	}

	return &recordpb.InsertResponse(), nil
}
