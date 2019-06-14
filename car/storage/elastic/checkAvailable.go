package elasticstore

import "context"

//CheckAvailable ...
func (estor *ElasticStore) CheckAvailable(ctx context.Context, cid string) (bool, error) {
	return estor.Lookup(ctx, cid)
}
