package car

import (
	"context"

	"github.com/md3756/ustart_tutorial/car/carpb"
)

// Pull retreives all car data based off of a car name
func (car *Car) Pull(ctx context.Context, req *carpb.PullRequest) (*carpb.PullResponse, error) {

	ca, err := car.strg.Lookup(ctx, req.UUID)
	if err != nil {
		return nil, err
	}

	return &carpb.PullResponse{CarProfile: &ca}, nil

}
