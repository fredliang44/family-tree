package middleware

import (
	"family-tree/db"
	"family-tree/utils"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// AuthMiddleware is a middleware to validate
var AuthMiddleware = &jwt.GinJWTMiddleware{
	Realm:      "Auth Middleware",
	Key:        []byte(utils.AppConfig.Server.SecretKey),
	Timeout:    time.Hour * 24,
	MaxRefresh: time.Hour,
	Authenticator: func(username string, password string, c *gin.Context) (string, bool) {
		res, err := db.FetchUserCache(username)
		if err != nil {
			log.Println("User Cache Do Not Exist", err)
			res, err = db.FetchUserFromMongo(username)
			if err != nil {
				log.Println("fetchUserFromMongo", err)
				return "User Do Not Exist", false
			}
		}
		if res.IsActivated != true {
			log.Println("GetUser: ", err)
			return "Please verify your account", false
		}
		isOK := CheckPasswordHash(password, res.Password)
		if isOK {
			cache, _ := msgpack.Marshal(&res)
			db.RedisClient.Set(res.Username, cache, 0)
			return res.Username, true
		}

		// Wrong passwork in cache, fetch user from mongo
		log.Println(username, "Wrong passwork in cache, fetch user from mongo")
		res, err = db.FetchUserFromMongo(username)
		if err != nil {
			log.Println("fetchUserFromMongo", err)
		}

		isOK = CheckPasswordHash(password, res.Password)

		if isOK {
			db.LoadUserCache(res)
			return res.Username, true
		}

		return username, false
	},
	Authorizator: func(username string, c *gin.Context) bool {

		res, err := db.FetchUserCache(username)

		if err != nil {
			log.Println("User Cache Do Not Exist", err)
			res, err = db.FetchUserFromMongo(username)
			if err != nil {
				log.Println("fetchUserFromMongo", err)
				return false
			}
		}

		if res.Username == "" {
			log.Println("GetUser: ", err)
			return false
		}

		return true
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	},
	// TokenLookup is a string in the form of "<source>:<name>" that is used
	// to extract token from the request.
	// Optional. Default value "header:Authorization".
	// Possible values:
	// - "header:<name>"
	// - "query:<name>"
	// - "cookie:<name>"

	TokenLookup: "header:Authorization",
	// TokenLookup: "query:token",
	// TokenLookup: "cookie:token",

	// TokenHeadName is a string in the header. Default value is "Bearer"
	TokenHeadName: "Bearer",

	// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
	TimeFunc: time.Now,
}

// CheckPasswordHash is a func to check password hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashPassword is a func to hash password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
