package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-17686.c302.asia-northeast1-1.gce.cloud.redislabs.com:17686",
		Password: "Xs5I7us2J8LTawcUcwSlf5o14urgiWlr",
		DB:       0,
	})

	fmt.Println(client)

	ctx := context.Background()
	// String
	// Set value
	err := client.Set(ctx, "car", "bmw", 0).Err()

	if err != nil {
		log.Fatalf("Error %v", err)
	}
	// Get value
	val, err := client.Get(ctx, "car").Result()

	if err != nil {
		log.Fatalf("Error %v", err)
	}
	fmt.Println("Val is", val)

	err = client.Del(ctx, "car").Err()

	// Delete value
	if err != nil {
		log.Fatalf("Error %v", err)
	} else {
		fmt.Println("Delete successfully")
	}

	// Hash
	keys := map[string]string{
		"company": "Fancy",
		"age":     "20",
	}
	for key, val := range keys {
		err := client.HSet(ctx, "test", key, val).Err()
		if err != nil {
			log.Fatalf("Error %v", err)
		}
	}

	res := client.HGetAll(ctx, "test").Val()
	fmt.Println(res)

	err = client.Del(ctx, "test").Err()

	if err != nil {
		log.Fatalf("Error %v", err)
	} else {
		fmt.Println("Delete successfully")
	}

	// Set
	err = client.SAdd(ctx, "uSet", "Khanh").Err()

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	ans := client.SMembers(ctx, "uSet").Val()

	fmt.Println(ans)

	err = client.SRem(ctx, "uSet", "Khanh").Err()

	if err != nil {
		log.Fatalf("Error %v", err)
	} else {
		fmt.Println("Delete successfully")
	}

	// Sorted set
	err = client.ZAdd(ctx, "sSet", redis.Z{
		Score:  2,
		Member: "Khanh",
	}).Err()

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	var val1 float64
	val1, err = client.ZScore(ctx, "sSet", "Khanh").Result()

	if err != nil {
		log.Fatalf("Error %v", err)
	}

	fmt.Println(val1)

	err = client.ZRem(ctx, "sSet", "Khanh").Err()

	if err != nil {
		log.Fatalf("Erorr %v", err)
	} else {
		fmt.Println("Delete successfully")
	}

	// hyperloglog
	err = client.PFAdd(ctx, "hello", "x1", "x2", "x3").Err()

	if err != nil {
		log.Fatalf("Erorr %v", err)
	}

	valx := client.PFCount(ctx, "hello").Val()

	fmt.Println(valx)

	// List
	err = client.RPush(ctx, "hola", "1", "2", "3", "4").Err()

	if err != nil {
		log.Fatalf("Erorr %v", err)
	}
	fmt.Println(client.LLen(ctx, "hola"))
	fmt.Println(client.LRange(ctx, "hola", 0, -1))
}
