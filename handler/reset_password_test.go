package handler

import (
	"family-tree/utils"
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGenResetCode(t *testing.T) {
	r := gofight.New()

	// check err sending code
	r.POST("/reset_password_code").
		SetDebug(true).
		SetJSON(gofight.D{
			"username": "test_err_sms1",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, `{"err_id":"", "message":"Testing err sms"}`, r.Body.String())
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})

	r.POST("/reset_password_code").
		SetDebug(true).
		SetJSON(gofight.D{
			"username": "test_family_tree",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, "{\"code\":200, \"message\":\"OK\"}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
	r.POST("/reset_password_code").
		SetDebug(true).
		SetJSON(gofight.D{
			"username": "test_fasdaamily_tree",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, `{"msg":"No Such User\n"}`, r.Body.String())
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
}

func TestResetPassword(t *testing.T) {
	r := gofight.New()
	r.POST("/reset").
		SetDebug(true).
		SetJSON(gofight.D{
			"username":   "123",
			"password":   utils.AppConfig.Root.Password,
			"verifyCode": "331",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, `{"code":400,"message":"Error occur when fetching user"}`, r.Body.String())
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})
	r.POST("/reset").
		SetDebug(true).
		SetJSON(gofight.D{
			"username":   utils.AppConfig.Root.Username,
			"password":   utils.AppConfig.Root.Password,
			"verifyCode": "",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, `{"code":406, "message":"Missing parameter"}`, r.Body.String())
			assert.Equal(t, http.StatusNotAcceptable, r.Code)
		})
	r.POST("/reset").
		SetDebug(true).
		SetJSON(gofight.D{
			"username":   utils.AppConfig.Root.Username,
			"password":   utils.AppConfig.Root.Password,
			"verifyCode": "233",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, `{"code":417, "message":"Wrong code"}`, r.Body.String())
			assert.Equal(t, http.StatusExpectationFailed, r.Code)
		})
	r.POST("/reset").
		SetDebug(true).
		SetJSON(gofight.D{
			"username":   utils.AppConfig.Root.Username,
			"password":   utils.AppConfig.Root.Password,
			"verifyCode": "2333",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, "{\"code\":200, \"message\":\"OK\"}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}
