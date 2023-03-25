package main

import (
	"fmt"
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
}
