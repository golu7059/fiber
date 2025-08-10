package main

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
    // Initialize a Redis client.
    client := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", // Redis server address
        DB:   0,                // Default DB
    })

    // Set a key in Redis.
    err := client.Set(ctx, "key", "value", 0).Err()
    if err != nil {
        panic(err)
    }

    // Retrieve the key from Redis.
    val, err := client.Get(ctx, "key").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("key:", val)

    // Output: key: value
}