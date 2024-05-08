package database

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"sync-finance/util"

	"github.com/redis/go-redis/v9"
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
