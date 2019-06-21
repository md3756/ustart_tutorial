package elasticstore

import (
	"context"
)

//ToggleAvailable ...
func (estor *ElasticStore) ToggleAvailable(ctx context.Context, cid string, avail bool) error {
	avail = !avail
	_, err := estor.client.Update().
		Index(estor.eIndex).
		Type(estor.eType).
		Id(cid).
		Doc(map[string]interface{}{"Available": avail}).
		Do(ctx)
	return err
}
