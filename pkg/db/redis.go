package db

import "github.com/jerrywange2025/newke_server_lib/pkg/redis"

func NewRedis(addr string, password string) *redis.Conn {
	return redis.New(addr, password)
}
