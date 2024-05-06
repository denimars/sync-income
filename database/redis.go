package database

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync-finance/util"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	ErrNil = errors.New("no matching record in redis database")
	Ctx    = context.Background()
)

func RedisConnection() (*redis.Client, error) {
	util.LoadEnv()
	port, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_URL"), os.Getenv("REDIS_PORT")),
		Password: "",
		DB:       port,
	})

	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}

	return client, nil
}

func MysqlConnection() *gorm.DB {
	util.LoadEnv()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DB_NAME"))
	fmt.Println(dsn)
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
