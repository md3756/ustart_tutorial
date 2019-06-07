package car

import "github.com/md3756/ustart_tutorial/car/storage"

// Config determines the runtime behavior of the Elastic-backed customer server
type Config struct {
	useDummy      bool
	StorageConfig *storage.Config
}
