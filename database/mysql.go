package database

import (
	"fmt"
	"log"
	"os"
	"sync-finance/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func MysqlConnection() *gorm.DB {
	util.LoadEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB_NAME"))
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}

func init() {
	db := MysqlConnection()
	err := db.AutoMigrate()
	if err != nil {
		panic(err)
	}
}
