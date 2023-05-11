package db

import (
	"os"

	"github.com/go-redis/redis"
)

func ConnRedis() (*redis.Client, error) {
	var clientRedis = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS"),
		// Addr:     "localhost:6379",
		Password: "", // password Redis Anda (jika ada)
		DB:       0,  // database Redis Anda
	})

	_, err := clientRedis.Ping().Result()
	if err != nil {
		return nil, err
	}

	return clientRedis, nil
}
