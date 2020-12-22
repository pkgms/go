/**
 * Created by zc on 2020/10/27.
 */
package database

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	page, _ := strconv.Atoi(
		r.URL.Query().Get("page"))
	size, _ := strconv.Atoi(
		r.URL.Query().Get("size"))
	return PaginateDirect(page, size)
}

func PaginateDirect(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page < 1 {
			page = 1
		}
		switch {
		case size > 100:
			size = 100
		case size < 1:
			size = 10
		}
		return db.Offset((page - 1) * size).Limit(size)
	}
}
