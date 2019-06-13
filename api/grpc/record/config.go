package recordapi

import (
	"github.com/md3756/ustart_tutorial/record"
)

// Config for auth api
type Config struct {
	RecordCfg *record.Config
	Port      int //Recomended 5101
}
