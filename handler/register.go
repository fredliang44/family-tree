package handler

import (
	"errors"
	"family-tree/db"
	"family-tree/middleware"
	"family-tree/utils"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/validator.v2"
)

// GenCode is a func to gen request code
func GenCode(c *gin.Context) {
	var info register
	c.BindJSON(&info)

	count, err := db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": info.Username}).Count()
	if count != 0 {
		err = errors.New("Username exists")
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": fmt.Sprintln(err), "code": http.StatusConflict})
		return
	}

	// log CreatedTime
	info.CreatedTime = time.Now()

	// gen verifyCode
	rand.Seed(time.Now().Unix())
	info.VerifyCode = fmt.Sprintf("%04d", rand.Intn(10000))

	// handle password
	if err := validator.Validate(info); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": err, "code": http.StatusNotAcceptable})
		return
	}

	info.Password, _ = middleware.HashPassword(info.Password)

	// send sms
	// Using tencent cloud
	// isOK, msg, errID := utils.SendDYSMS(data.Phone, "SMS_133979618", `{"code":"`+code+`"}`) DAYU example
	isOK, msg, errID := utils.SendQCSMS(info.Phone, 96385, []string{"Family Tree", info.VerifyCode})

	if !isOK {
		c.JSON(http.StatusBadRequest, gin.H{"message": msg, "err_id": errID})
		return
	}

	// save user info
	info.ID = ai.Next("user")
	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Insert(info)

	c.JSON(http.StatusOK, gin.H{"message": "OK", "code": http.StatusOK})
}

// RegisterHandler is a func to verify sms code
func RegisterHandler(c *gin.Context) {
	var data register
	var info register
	c.BindJSON(&data)

	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Find(bson.M{"username": data.Username}).One(&info)

	if info.Username == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "No such user", "code": http.StatusNotFound})
		return
	}

	if info.IsActivated == true && info.VerifyCode == data.VerifyCode {
		c.JSON(http.StatusOK, gin.H{"message": "OK", "status": "Already Verifyed", "code": http.StatusOK})
		return
	}

	fmt.Println("info.VerifyCode: ", info.VerifyCode, "data.VerifyCode: ", data.VerifyCode)

	if info.VerifyCode == data.VerifyCode {
		info.IsActivated = true
		db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Update(bson.M{"username": data.Username}, &info)
		c.JSON(http.StatusOK, gin.H{"message": "OK", "status": "Verifyed", "code": http.StatusOK})
		return
	}
	c.JSON(http.StatusForbidden, gin.H{"message": "Wrong Verify Code", "code": http.StatusForbidden})
}

type register struct {
	ID          uint64    `json:"id" bson:"_id,omitempty"`
	Username    string    `validate:"min=3,max=20" bson:"username" json:"username"`
	Password    string    `validate:"min=3,max=30,regexp=^((?=.*\d)|(?=.*\W+))(?![.\n])(?=.*[a-z]).*$" bson:"password" json:"password"`
	Phone       string    `validate:"len=11,regexp=^1[34578]\d{9}$" bson:"phone" json:"phone"`
	VerifyCode  string    `bson:"verifyCode" json:"verifyCode"`
	IsActivated bool      `bson:"isActivated" json:"isActivated"`
	CreatedTime time.Time `bson:"createdTime" json:"createdTime"`
}
