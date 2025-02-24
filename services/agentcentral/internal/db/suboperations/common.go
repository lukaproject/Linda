package suboperations

import (
	"Linda/baselibs/abstractions"

	"gorm.io/gorm"
)

func listQueryAsync[T any](
	responsesChan chan *T,
	listQueryPacker abstractions.ListQueryPacker,
	dbi *gorm.DB,
	primaryKey string,
) {
	defer close(responsesChan)
	dbCurrent := listQueryPacker.PackListQuery(primaryKey, dbi.Model(new(T)))
	rows, err := dbCurrent.Rows()
	if err != nil {
		return
	}
	for rows.Next() {
		var entity = new(T)
		if err = dbCurrent.ScanRows(rows, entity); err != nil {
			break
		}
		responsesChan <- entity
	}
}
