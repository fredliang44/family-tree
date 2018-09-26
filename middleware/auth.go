package middleware

import (
	"github.com/appleboy/gin-jwt"
	"github.com/fredliang44/family-tree/db"
	"github.com/fredliang44/family-tree/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-errors/errors"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

var AuthMiddleware = getAuthMiddleware()

func getAuthMiddleware() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "Auth Middleware",
		Key:         []byte(utils.AppConfig.Server.SecretKey),
		Timeout:     refreshTimeOut(),
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Username: claims["id"].(string),
			}
		},

		// Authenticator for login usage
		Authenticator: authLogin,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			v, ok := data.(*User)

			if ok {

				username := v.Username

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
			}

			return false
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
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal(err)
	}
	return authMiddleware
}

// User demo
type User struct {
	Username string
	IsAdmin  bool
}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

// @Summary Refresh Token
// @Description Refresh Token
// @Tags additional
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} utils.TokenResp
// @Failure 400 {object} utils.ErrResp
// @Router /refresh_token [get]
func refreshTimeOut() time.Duration {
	return time.Hour * 24
}

// @Summary Login
// @Description Login
// @Tags Login
// @Accept  json
// @Produce  json
// @Param 	Login body utils.LoginReq true "Login"
// @Success 200 {object} utils.TokenResp
// @Failure 400 {object} utils.ErrResp
// @Router /login [post]
func authLogin(c *gin.Context) (interface{}, error) {
	var loginVals login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	username := loginVals.Username
	password := loginVals.Password

	// Wrong password in cache, fetch user from mongo
	log.Println(username, "Wrong password in cache, fetch user from mongo")
	res, err := db.FetchUserFromMongo(username)
	if err != nil {
		log.Println("fetchUserFromMongo", err)
	}

	if res.IsActivated != true {
		log.Println("IsActivated: ", err)
		return nil, errors.New("Please verify your account")
	}

	if res.IsValidated != true {
		log.Println("IsValidated: ", err)
		return nil, errors.New("Please contact admin to validate your account")
	}

	isOK := CheckPasswordHash(password, res.Password)

	if isOK {
		db.LoadUserCache(res)
		return &User{
			Username: username,
			IsAdmin:  res.IsAdmin,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
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
