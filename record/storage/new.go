package storage

import (
	sqlstore "github.com/md3756/ustart_tutorial/record/storage/sql"
)

// NewSQL determines the runtime behavior of the ElasticSearch-backed customer server
func NewSQL(config *Config) (Storage, error) {
	strg, err := &Config{SQLConfig: sqlstore.NewConfig()}
	return strg, err
}
