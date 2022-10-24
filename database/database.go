package database

import (
	"fmt"
	"go-wms-poc/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(configuration config.Configuration) *gorm.DB {
	dbConfig := configuration

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.DbUsername,
		dbConfig.DbPassword,
		dbConfig.DbHost,
		dbConfig.DbPort,
		dbConfig.DbDatabase,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("error connecting database")
	}

	return db
}
