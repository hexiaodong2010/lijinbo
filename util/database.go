package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"lijinbo/config"
)

func GetMySqlDb() *gorm.DB {
	db, err := gorm.Open("mysql", config.MysqlDsn)
	if err !=nil {
		panic(err)
	}
	if config.RunMode != config.ProEnv {
		db.LogMode(true)
	}else {
		db.LogMode(false)
	}
	return db
}