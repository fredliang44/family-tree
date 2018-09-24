package db

import (
	t "github.com/fredliang44/family-tree/graphql/types"
	"github.com/fredliang44/family-tree/utils"
	"log"
	"time"

	"github.com/go-redis/redis"
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
	//cache, _ := msgpack.Marshal(&user)
	CacheClient.Set(user.Username, user, time.Hour*24)
}

// FetchUserCache is a func to fetch user data from redis
func FetchUserCache(username string) (user t.User, err error) {
	var res t.User
	raw, isExist := CacheClient.Get(username)
	if !isExist {
		log.Println("Get User from cache: ", err)
		res, err = FetchUserFromMongo(username)
		return res, err
	}

	return raw.(t.User), nil
}

// CheckAdminByUsername is a func to Check Admin By Username
func CheckAdminByUsername(username string) (isAdmin bool) {

	cacheIsAdmin, isExist := CacheClient.Get("isAdmin" + username)

	if !isExist {
		res := checkAdminFromMongo(username)
		CacheClient.Set("isAdmin"+username, isAdmin, time.Minute)
		return res
	}

	return cacheIsAdmin.(bool)
}

// FetchUserIDByUsername is a func to Fetch UserID By Username
func FetchUserIDByUsername(username string) (userID uint64, err error) {
	cacheUserID, isExist := CacheClient.Get("userID" + username)

	if !isExist {
		userID, err := fetchUserIDFromMongo(username)
		CacheClient.Set("userID"+username, userID, time.Hour*24)
		return userID, err
	}

	return cacheUserID.(uint64), nil
}
