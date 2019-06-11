package elasticstore

import (
	"context"

	"github.com/md3756/ustart_tutorial/record/recordpb"
)

//Register creates a new ES document for a new registering record
func (estor *ElasticStore) Register(ctx context.Context, carID string, userID string) error {

	//Lock just to make sure no two people can sign up with the same username at the same time
	newRecordLock.Lock()
	defer newRecordLock.Unlock()

	_, err := estor.client.Index().
		Index(estor.eIndex).
		Type(estor.eType).
		BodyJson(recordpb.Record{
			CarID:  carID,
			UserID: userID,
		}).
		Id(carID).
		Do(ctx)

	return err
}
