package db

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type dbclient struct {
	Redis *redis.Client
}

var ctx = context.Background()

func NewClient() (*dbclient, error) {
	rdbClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Test connection
	err := rdbClient.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return &dbclient{
		Redis: rdbClient,
	}, nil
}

// GetKey is wrapper function for Get Redis Command
func (db *dbclient) GetKey(key string) (string, error) {
	val, err := db.Redis.Get(ctx, key).Result()
	return val, err
}
