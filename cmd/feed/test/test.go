package main

import (
	"context"
	"fmt"
	"miniTikTok/cmd/favorite/dal/redisDb"
)

var ctx = context.Background()

func main() {
	redisDb.Init()
	n, err := redisDb.RDB.Exists(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	if n > 0 {
		fmt.Println("存在")
	} else {
		fmt.Println("不存在")
	}
}
