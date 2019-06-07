package car

import (
	"github.com/md3756/ustart_tutorial/car/storage"
)

//Car is an implementation of the car service as defined in service.proto
type Car struct {
	strg          storage.Storage
	defaultAvatar string
	defaultBanner string
}

// New returns a new Eclient customer service
func New(cfg *Config) (*Car, error) {
	// if cfg.useDummy

	storg, err := storage.NewES(cfg.StorageConfig)

	ca := &Car{
		strg: storg,
	}

	return ca, err
}
