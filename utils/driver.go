package utils

import (
	"api-jasa-pengiriman/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseDriver string

const (
	Mysql DatabaseDriver = "mysql"
)

type DatabaseConnection struct {
	Driver DatabaseDriver

	Mysql *gorm.DB
}

func NewConnectionDatabase(config *config.AppConfig) *DatabaseConnection {
	var db DatabaseConnection
	db.Driver = Mysql
	db.Mysql = newMysql(config)

	return &db
}

func newMysql(config *config.AppConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Database.DBURL), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func (db *DatabaseConnection) CloseConnection() {
	if db.Mysql != nil {
		db, _ := db.Mysql.DB()
		db.Close()
	}
}
