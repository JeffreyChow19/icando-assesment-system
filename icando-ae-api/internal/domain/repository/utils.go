package repository

import (
	"fmt"
	"gorm.io/gorm"
)

func Paginate(query *gorm.DB, page int, limit int) {
	if page < 0 {
		page = 1
	}
	if limit < 0 {
		limit = 10
	}
	offset := (page - 1) * limit
	query.Offset(offset)
}

func Sort(query *gorm.DB, asc bool, sortBy string) {
	ascKey := "asc"

	if !asc {
		ascKey = "desc"
	}
	query.Order(fmt.Sprintf("%s %s", sortBy, ascKey))
}
