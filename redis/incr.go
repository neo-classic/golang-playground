package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{"localhost:6379"},
		DB:    2,
	})

	pipe := client.Pipeline()
	incr := pipe.IncrBy(ctx, "qweqwe", int64(-2))
	pipe.Expire(ctx, "qweqwe", 300*time.Second)

	_, err := pipe.Exec(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Printf("incr val = %d \n", incr.Val())
}
