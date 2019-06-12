package storage

import (
	elasticstore "github.com/md3756/ustart_tutorial/record/storage/elastic"
)

// NewES determines the runtime behavior of the ElasticSearch-backed customer server
func NewES(config *Config) (Storage, error) {
	return elasticstore.New(config.ElasticConfig)
}
