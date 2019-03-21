package main

import (
    "github.com/go-redis/redis"
    "fmt"
)

func r_redis_init() {
    redisdb := redis.NewClient(&redis.Options{
        Addr: ":6379",
    })

    redisdb.WrapProcess(func(old func(cmd redis.Cmder) error) func(cmd redis.Cmder) error {
        return func(cmd redis.Cmder) error {
            fmt.Printf("starting processing: <%s>\n", cmd)
            err := old(cmd)
            fmt.Printf("finished processing: <%s>\n", cmd)
            return err
        }
    })

    redisdb.Ping()
}