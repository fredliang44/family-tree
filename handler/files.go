package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fredliang44/family-tree/db"
	"github.com/fredliang44/family-tree/utils"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

// GetPolicyTokenHandler is a func to get PolicyTokenHandler for OSS file upload
// @Summary Get Policy Token
// @Description Get Policy Token, for details: https://help.aliyun.com/document_detail/31926.html?spm=a2c4g.11186623.6.635.cscSyI
// @Tags files
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param table query string false "table (user, group, project, company)" Enums(user, group, project, company)
// @Param table_id query int false "tableid: int type " mininum(1)
// @Param field query string false "field in table (avatar, images, logo)" Enums(avatar, images, logo)
// @Param action query string false "action (init, add)" Enums(init, add)
// @Success 200 {object} utils.PolicyToken
// @Router /files/token [get]
func GetPolicyTokenHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c) // load from jwt middleware
	var username string
	if claims["id"] == nil {
		username = "temp"
	} else {
		username = claims["id"].(string)
	}

	fileName := c.Query("file_name")

	fmt.Println("filename", fileName)
	table := c.Query("table")
	field := c.Query("field")
	action := c.Query("action")
	tableID, err := strconv.Atoi(c.Query("table_id"))

	if err != nil {
		c.JSON(http.StatusConflict, utils.ErrResp{Message: fmt.Sprintln(err), Code: http.StatusConflict})
		return
	}

	c.JSON(http.StatusOK, utils.GetPolicyToken(username, table, field, action, tableID, fileName))
}

// FilesCallBackHandler is a func to handle call back request
// @Summary Call Back
// @Description  Call Back
// @Tags files
// @Accept  json
// @Produce  json
// @Param 	Register body utils.CallBackBody true "Callback Param"
// @Success 200 {object} utils.PolicyToken
// @Router /files/callback [get]
func FilesCallBackHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c) // load from jwt middleware
	var username string
	if claims["id"] == nil {
		username = "temp"
	} else {
		username = claims["id"].(string)
	}

	var data utils.CallBackBody

	buf, err := ioutil.ReadAll(c.Request.Body)
	buf = []byte(strings.Replace(string(buf), "\"\"", "\"", -1))
	fmt.Println(string(buf))
	err = json.Unmarshal(buf, &data)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, utils.ErrResp{Code: http.StatusNotAcceptable, Message: err.Error()})
		log.Print("bodyErr ", err.Error())
		return
	}

	action := data.Action

	switch action {
	case "init":
		action = "$update"
		break
	case "add":
		action = "$push"
		break
	default:
		c.JSON(http.StatusNotAcceptable, utils.ErrResp{Code: http.StatusNotAcceptable, Message: "Action not found"})
		return
	}

	fmt.Println(data)

	if data.Field == "avatar" || data.Field == "logo" {
		err = db.DBSession.DB(utils.AppConfig.Mongo.DB).C(data.Table).Update(
			bson.M{"_id": data.TableID},
			bson.M{
				"$set": bson.M{
					data.Field: data.FilePath + data.FileName,
				},
			})
	} else {
		err = db.DBSession.DB(utils.AppConfig.Mongo.DB).C(data.Table).Update(
			bson.M{"username": username},
			bson.M{
				action: bson.M{
					data.Field: []string{data.FilePath},
				},
			})
	}

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusNotFound, utils.ErrResp{Code: http.StatusNotFound, Message: err.Error()})
		return
	}

	// fix ali oss handler error
	rdr := bytes.NewBuffer(buf).String()
	rdr = strings.Replace(rdr, "\"\"", "\"", -1)

	json.Unmarshal([]byte(rdr), &data)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": data, "message": "OK"})
}
