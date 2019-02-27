package NtmRedis

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

type BmRedis struct {
	Host     string
	Port     string
	Password string
	Database string
}

func (r BmRedis) NewRedisDBDaemon(args map[string]string) *BmRedis {
	return &BmRedis{
		Host:     args["host"],
		Port:     args["port"],
		Password: args["password"],
		Database: args["database"]}
}

func (r BmRedis) GetRedisClient() *redis.Client {

	db, _ := strconv.Atoi(r.Database)
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprint(r.Host, ":", r.Port),
		Password: r.Password, // no password set
		DB:       db,         // use default DB
	})
	return client
}

func (r BmRedis) CheckToken(token string) error {
	client := r.GetRedisClient()
	defer client.Close()

	_, err := client.Get(token).Result()
	if err == redis.Nil {
		return errors.New("token not exist")
	} else if err != nil {
		fmt.Println(err.Error())
		panic(err)
		return err
	} else {
		return nil
	}
}

func (r BmRedis) PushToken(token string, exptime time.Duration) error {
	client := r.GetRedisClient()
	defer client.Close()

	pipe := client.Pipeline()

	pipe.Incr(token)
	pipe.Expire(token, exptime)

	_, err := pipe.Exec()
	fmt.Println(token)
	return err
}
