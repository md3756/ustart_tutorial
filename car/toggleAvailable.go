package car

import (
	"context"

	"github.com/md3756/ustart_tutorial/car/carpb"
)

// ToggleAvailable ...
func (car *Car) ToggleAvailable(ctx context.Context, req *carpb.ToggleRequest) (*carpb.ToggleResponse, error) {
	c, err := car.strg.CheckAvailable(ctx, req.CID)
	if err != nil {
		return nil, err
	}
	available, er := car.strg.UpdateAvailable(ctx, req.CID, c)

	return &carpb.ToggleResponse{NewStatus: available}, er
}
