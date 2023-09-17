package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"test.com/project-common/config"
	"time"
)

var Rc *RedisCache

type RedisCache struct {
	rdb *redis.Client
}

func init() {
	// 做redischche连接server操作
	config.InitConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.C.RedisConfig.Address,
		Password: config.C.RedisConfig.Password,
		DB:       config.C.RedisConfig.Db,
	})
	Rc = &RedisCache{
		rdb: rdb,
	}
}

func (*RedisCache) Put(ctx context.Context, key, value string, expire time.Duration) error {
	err := Rc.rdb.Set(ctx, key, value, expire).Err()
	return err
}

func (*RedisCache) Get(ctx context.Context, key string) (string, error) {
	result, err := Rc.rdb.Get(ctx, key).Result()
	return result, err
}
