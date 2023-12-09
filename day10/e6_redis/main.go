package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8" // 注意导入的是新版本
	"time"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}

func V8Example() {
	ctx := context.Background()
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output:
	// key value
	// key2 does not exist
}

func redisExample() {
	ctx := context.Background()
	zsetKey := "language_rank"
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C++"},
	}
	// ZAdd
	num, err := rdb.ZAdd(ctx, zsetKey, languages...).Result()
	if err != nil {
		fmt.Println("ZAdd failed, err:", err)
		return
	}
	fmt.Printf("ZAdd %d success.\n", num)

	// 把Golang的分数加10
	newScore, err := rdb.ZIncrBy(ctx, zsetKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Println("ZIncrBy failed, err:", err)
		return
	}
	fmt.Printf("Golang's score is %.1f now.\n", newScore)

	// 取分数最高的3个
	result, err := rdb.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Println("ZRevRange failed, err:", err)
		return
	}
	for _, z := range result {
		fmt.Println(z.Member, z.Score)
	}

	// 取95~100分的
	opt := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	result, err = rdb.ZRangeByScoreWithScores(ctx, zsetKey, opt).Result()
	if err != nil {
		fmt.Println("ZRangeByScore failed, err:", err)
		return
	}
	for _, z := range result {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	err := initClient()
	if err != nil {
		fmt.Println("connect redis failed, err:", err)
		return
	}
	fmt.Println("connect redis success!")

	V8Example()
	redisExample()
}
