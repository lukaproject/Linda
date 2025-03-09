package abstractions

import (
	"errors"
	"net/url"
	"strconv"

	"github.com/lukaproject/xerr"
	"gorm.io/gorm"
)

type ListQueryPacker interface {
	PackListQuery(primaryKey string, db *gorm.DB) *gorm.DB
}

type listQueryPacker struct {
	prefix      *string
	createAfter *int64
	limit       *int
	idAfter     *string
	// an external prarameter for multifields.
	_ *map[string]string
}

func (lqp *listQueryPacker) PackListQuery(primaryKey string, db *gorm.DB) *gorm.DB {
	dbCurrent := db
	if lqp.prefix != nil {
		dbCurrent = dbCurrent.Where(primaryKey+" LIKE ?", *lqp.prefix+"%")
	}
	if lqp.idAfter != nil {
		dbCurrent = dbCurrent.Where(primaryKey+" >= ?", *lqp.idAfter)
	}
	if lqp.createAfter != nil {
		dbCurrent = dbCurrent.Where("create_time_ms >= ?", *lqp.createAfter)
	}
	if lqp.limit != nil {
		dbCurrent = dbCurrent.Limit(*lqp.limit)
	}
	return dbCurrent
}

func NewListQueryPacker(
	query url.Values,
) (ListQueryPacker, error) {

	if query.Has("createAfter") && query.Has("idAfter") {
		return nil, errors.New("createAfter and idAfter couldn't in same packer")
	}
	lqp := &listQueryPacker{}
	if query.Has("prefix") {
		prefix := query.Get("prefix")
		lqp.prefix = &prefix
	}

	if query.Has("createAfter") {
		createAfter := xerr.Must(strconv.ParseInt(query.Get("createAfter"), 10, 64))
		lqp.createAfter = &createAfter
	}

	if query.Has("idAfter") {

		idAfter := query.Get("idAfter")
		lqp.idAfter = &idAfter
	}

	if query.Has("limit") {
		limit := xerr.Must(strconv.Atoi(query.Get("limit")))
		lqp.limit = &limit
	}

	return lqp, nil
}
