package utils

import (
	"fmt"
	"gorm.io/gorm"
)

func QueryPaginate(db *gorm.DB, page int, limit int) *gorm.DB {
	offset := (page - 1) * limit
	return db.Offset(offset).Limit(limit)
}

func QuerySortBy(db *gorm.DB, sortBy string, asc bool) *gorm.DB {
	ascKey := "asc"

	if !asc {
		ascKey = "desc"
	}

	return db.Order(fmt.Sprintf("%s %s", sortBy, ascKey))
}
