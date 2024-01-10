package dbutil

import (
	"context"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

// TODO
func GormDB(ctx context.Context) (*gorm.DB, error) {

	if db == nil {
		var err error
		db, err = gorm.Open(mysql.Open("root:@tcp(localhost:3306)/factory_db?parseTime=true&multiStatements=true"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		return db, err
	}
	return db.Debug(), nil

}
