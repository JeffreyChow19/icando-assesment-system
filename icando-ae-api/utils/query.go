package utils

import (
	"fmt"
	"gorm.io/gorm"
)

func QueryPaginate(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * limit
		return db.Offset(offset).Limit(limit)
	}
}

func QuerySortBy(sortBy string, asc bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		ascKey := "asc"

		if !asc {
			ascKey = "desc"
		}

		return db.Order(fmt.Sprintf("%s %s", sortBy, ascKey))
	}
}
