package main

import (
	"context"
	"fmt"
	"github.com/angelmotta/go-flow/db"
	"github.com/redis/go-redis/v9"
	"log"
)

func Check(err error) {
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Golang with in-memory database")
	ctx := context.Background()

	dbClient, errDb := db.NewClient()
	Check(errDb)

	// Using Set command from the client connection
	Check(dbClient.Redis.Set(ctx, "mykey", "myval", 0).Err())

	// Using Get command from wrapper function (method from db struct)
	val, err := dbClient.Redis.Get(ctx, "mykey").Result()
	Check(err)
	fmt.Println("mykey:", val)

	// Get null value
	val2, err := dbClient.Redis.Get(ctx, "key2").Result()
	Check(err)
	fmt.Println("key2:", val2)

	// Set value
	Check(dbClient.Redis.Set(ctx, "sol-dollar", 3.793, 0).Err())

	// Get value
	val3, err := dbClient.Redis.Get(ctx, "sol-dollar").Result()
	Check(err)
	fmt.Println("sol-dollar", val3)
}
