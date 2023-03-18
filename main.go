package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	fmt.Println("Golang with in-memory database")

	rdbClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	//pong, err := rdbClient.Ping(ctx).Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(pong)

	err := rdbClient.Set(ctx, "mykey", "myval", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdbClient.Get(ctx, "mykey").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("mykey", val)

	val2, err := rdbClient.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	err2 := rdbClient.Set(ctx, "sol-dollar", 3.793, 0).Err()
	if err2 != nil {
		panic(err)
	}

	val3, err := rdbClient.Get(ctx, "sol-dollar").Result()
	if err == redis.Nil {
		fmt.Println("sol-dollar key does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("sol-dollar", val3)
	}
}
