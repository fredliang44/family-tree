package handler

import (
	"family-tree/db"
	t "family-tree/graphql/types"
	m "family-tree/middleware"
	"family-tree/utils"
	"github.com/gin-gonic/gin"
	"github.com/night-codes/mgo-ai"
	"net/http"
	"time"
)

// @Summary InitDB
// @Description InitDB for unit test
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} utils.StdResp
// @Router /init_db [get]
// InitDB is a func to init db for unit test
func InitDB(c *gin.Context) {
	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").RemoveAll(nil)

	user := t.User{
		ID:          ai.Next("user"),
		Username:    utils.AppConfig.Root.Username,
		IsActivated: false,
		IsAdmin:     true,
		VerifyCode:  "2333",
		Phone:       "17777766666",
		CreatedTime: time.Now(),
	}

	user.Password, _ = m.HashPassword(utils.AppConfig.Root.Password)
	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Insert(user)

	user.Username = "test_err_sms1"
	user.Phone = "17777766667"
	user.ID = ai.Next("user")
	db.DBSession.DB(utils.AppConfig.Mongo.DB).C("user").Insert(user)

	c.JSON(http.StatusOK, utils.StdResp{Message: "OK", Code: http.StatusOK})
}
