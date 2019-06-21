package car

import (
	"context"

	"github.com/md3756/ustart_tutorial/car/carpb"
)

// ToggleAvailable ...
func (car *Car) ToggleAvailable(ctx context.Context, req *carpb.ToggleRequest) (*carpb.ToggleResponse, error) {
	bool c = CheckAvailable(req)
	return UpdateAvailable(c)
}