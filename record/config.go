package record

import "github.com/md3756/ustart_tutorial/record/storage"

// Config determines the runtime behavior of the Elastic-backed customer server
type Config struct {
	StorageConfig *storage.Config
}
