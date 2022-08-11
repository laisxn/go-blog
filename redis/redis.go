package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-gin/config"
	"time"
)

var rdb *redis.Client //创建redis

func init() {
	rdb = redis.NewClient(&redis.Options{
		//	端口需要改，这里是docker的端口
		Addr:         fmt.Sprintf("%s:%s", config.Get("redis.host"), config.Get("redis.port")),
		Password:     config.Get("redis.password"),      // no password set
		DB:           config.GetToInt("redis.database"), // use default DB
		PoolSize:     15,
		MinIdleConns: 10, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。
		//超时
		//DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		//ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		//WriteTimeout: 3 * time.Second, //写超时，默认等于读超时
		//PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
	})
}

func Client() *redis.Client {
	return rdb
}

const (
	AddCommentLockKey = "comment_lock_key:%s"
)

func CreateLock(key string) bool {
	var ctx = context.Background()

	set, err := rdb.SetNX(ctx, key, "1", 5*time.Second).Result()
	if err != nil {
		panic(err)
	}

	return set
}

func GetLock(key string) (string, error) {
	var ctx = context.Background()
	return rdb.Get(ctx, key).Result()
}
