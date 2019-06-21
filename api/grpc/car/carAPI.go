package carapi

import (
	"strconv"

	"github.com/md3756/ustart_tutorial/car"
)

//GPRCAPI is the GRPC API implementation for car
type GPRCAPI struct {
	prof *car.Car
	port string
}

// New creates a new car api, given the config
func New(cfg *Config) (*GPRCAPI, error) {
	prof, err := car.New(cfg.CarCfg)
	if err != nil {
		return nil, err
	}

	return &GPRCAPI{
		prof: prof,
		port: strconv.Itoa(cfg.Port),
	}, nil
}
