package car

import (
	"github.com/md3756/ustart_tutorial/car"
)

// RESTAPI implements a REST api
// as a wrapper around the car package.
type RESTAPI struct {
	car *car.Car
}

// New creates a new car api, given the config
func New(cfg *Config) (*RESTAPI, error) {
	car, err := car.New(cfg.CarCfg)
	if err != nil {
		return nil, err
	}

	return &RESTAPI{
		car: car,
	}, nil
}
