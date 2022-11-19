package adapters

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLAdapter struct {
	ConnectionString string
	DB               *gorm.DB
}

func newSqlContext(dsn string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	return
}
