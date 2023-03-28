package exchangestore

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type ExchangeStore struct {
	redis *redis.Client
}

var ctx = context.Background()

func New() (*ExchangeStore, error) {
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

	return &ExchangeStore{
		redis: rdbClient,
	}, nil
}

// GetExchange retrieves a Currency Exchange from the Store layer
func (db *ExchangeStore) GetExchange(key string) (string, error) {
	val, err := db.redis.Get(ctx, key).Result()
	return val, err
}
