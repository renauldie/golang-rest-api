package config

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

func LoadRedisConfig() *RedisConfig {
	dbNum, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("REDIS_PORT")
	if port == "" {
		port = "6379"
	}

	return &RedisConfig{
		RedisHost:     host,
		RedisPort:     port,
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisDB:       dbNum,
	}
}

func (c *RedisConfig) InitRedis() (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", c.RedisHost, c.RedisPort)

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: c.RedisPassword,
		DB:       c.RedisDB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
