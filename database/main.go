package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open(dbPath string) (*gorm.DB, error) {
	if db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{}); err != nil {
		return db, err
	} else {
		return db, nil
	}
}
