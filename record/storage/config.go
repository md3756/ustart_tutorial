package storage

import (
	//when i save the file, this import pops up
	"github.com/md3756/ustart_tutorial/car/storage"
	sqlstore "github.com/md3756/ustart_tutorial/record/storage/sql"
)

// Config determines the runtime behavior of the an either SQL or ElasticSearch backed customer server
type Config struct {
	// ElasticConfig *elasticstore.Config
	StorageConfig *storage.Storage
	SQLConfig     *sqlstore.Config
}

// ESNewConfig returns a default config object
func ESNewConfig() *Config {
	return &Config{SQLConfig: sqlstore.NewConfig()}
}

// SQLNewConfig returns a default config object
// func SQLNewConfig() *Config {
// 	return &Config{SQLConfig: sqlstore.NewConfig()}
// }
