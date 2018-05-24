package db

import (
	t "family-tree/graphql/types"
	"family-tree/utils"
	"github.com/getsentry/raven-go"
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
	"log"
)

// RedisClient locate mongo db client
var RedisClient = redisClient()

func redisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     utils.AppConfig.Redis.Host + ":" + utils.AppConfig.Redis.Port,
		Password: utils.AppConfig.Redis.Password,
		DB:       utils.AppConfig.Redis.DB,
	})
	return client
}

// LoadUserCache is a func to load user data to redis
func LoadUserCache(user t.User) {
	cache, _ := msgpack.Marshal(&user)
	RedisClient.Set(user.Username, cache, 0)
}

// FetchUserCache is a func to fetch user data from redis
func FetchUserCache(username string) (user t.User, err error) {
	var res = t.User{}
	resp, _ := RedisClient.Get(username).Bytes()
	err = msgpack.Unmarshal(resp, &res)
	if err != nil {
		raven.CaptureError(err, nil)
		log.Println("Get User from redis: ", err)
		res, err = FetchUserFromMongo(username)
		return res, err
	}
	return res, nil
}
