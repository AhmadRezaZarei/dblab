package dbutil

import (
	"context"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// TODO
func GormDB(ctx context.Context) (*gorm.DB, error) {

	if db == nil {
		var err error
		db, err = gorm.Open(mysql.Open(os.Getenv("DB_CONNECTION_STRING")), &gorm.Config{})
		return db, err
	}
	return db.Debug(), nil

}
