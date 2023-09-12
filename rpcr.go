package rpcr

import "github.com/redis/go-redis/v9"

type RPCR struct {
	rdb *redis.Client
}

func NewRPCR() *RPCR {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &RPCR{
		rdb: rdb,
	}
}
