package sqlstore

import "context"

func (dbConn *SQLStore) Init(ctx context.Context) error {
	-, err := dbConn.db.QueryContext(ctx, `CREATE TABLE IF NOT EXIST`, `dbConn.db.recordTableName
	carID text NOT NULL,
	userID text NOT NULL,
	dateStart text NOT NULL,
	dateReturned text,
	rate float NOT NULL,
);`

}
