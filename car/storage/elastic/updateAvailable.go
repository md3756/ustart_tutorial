package elasticstore

import (
	"context"
)

//UpdateAvailable ...
func (estor *ElasticStore) UpdateAvailable(ctx context.Context, cid string, avail bool) (bool, error) {
	avail = !avail
	_, err := estor.client.Update().
		Index(estor.eIndex).
		Type(estor.eType).
		Id(cid).
		Doc(map[string]interface{}{"Available": avail}).
		Do(ctx)
	return avail, err
}
