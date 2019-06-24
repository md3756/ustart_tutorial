package car

import (
	"context"

	"github.com/md3756/ustart_tutorial/car/carpb"
)

// ToggleAvailable ...
func (car *Car) ToggleAvailable(ctx context.Context, req *carpb.ToggleRequest) (*carpb.ToggleResponse, error) {
	c, err := car.strg.CheckAvailable(req)
	return car.strg.UpdateAvailable(c)
}
