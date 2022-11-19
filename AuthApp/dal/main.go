package dal

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLAdapter struct {
	ConnectionString string
	DB               *gorm.DB
}

func NewSqlContext(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	return
}
