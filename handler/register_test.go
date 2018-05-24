package handler

import (
	"family-tree/utils"
	"net/http"
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestGenCode(t *testing.T) {
	r := gofight.New()

	// check err sending code
	r.POST("/register_code").
		SetDebug(true).
		SetJSON(gofight.D{
			"username": "test_err_sms",
			"password": "test",
			"phone":    "17777766667",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, `{"err_id":"", "message":"Testing err sms"}`, r.Body.String())
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})

	r.POST("/register_code").
		SetDebug(true).
		SetJSON(gofight.D{
			"username": "uniquestudio",
			"password": "test",
			"phone":    "17777766666",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, `{"code":409, "msg":"Username exists\n"}`, r.Body.String())
			assert.Equal(t, http.StatusBadRequest, r.Code)
		})

	r.POST("/register_code").
		SetDebug(true).
		SetJSON(gofight.D{
			"username": "test_family_tree",
			"password": "test",
			"phone":    "17777766666",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, `{"code":200, "message":"OK"}`, r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})

	r.POST("/register_code").
		SetDebug(true).
		SetJSON(gofight.D{
			"username": "test_fasdasdasdasdadsasdasdhoahduhaoushoaamily_tree",
			"password": "test",
			"phone":    "17777766666",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.Equal(t, http.StatusNotAcceptable, r.Code)
		})
}

func TestRegister(t *testing.T) {
	r := gofight.New()

	r.POST("/register").
		SetDebug(true).
		SetJSON(gofight.D{
			"username":   utils.AppConfig.Root.Username,
			"verifyCode": "2333",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, "{\"code\":200,\"message\":\"OK\",\"status\":\"Verifyed\"}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
	r.POST("/register").
		SetDebug(true).
		SetJSON(gofight.D{
			"username":   "",
			"verifyCode": "2333",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, `{"message":"No such user", "code":404}`, r.Body.String())
			assert.Equal(t, http.StatusNotFound, r.Code)
		})
	r.POST("/register").
		SetDebug(true).
		SetJSON(gofight.D{
			"username":   utils.AppConfig.Root.Username,
			"verifyCode": "2333",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, "{\"code\":200,\"message\":\"OK\",\"status\":\"Already Verifyed\"}", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
	r.POST("/register").
		SetDebug(true).
		SetJSON(gofight.D{
			"username":   utils.AppConfig.Root.Username,
			"verifyCode": "2324",
		}).
		Run(GinEngine(), func(r gofight.HTTPResponse, rq gofight.HTTPRequest) {
			assert.JSONEq(t, "{\"message\":\"Wrong Verify Code\", \"code\":403}", r.Body.String())
			assert.Equal(t, http.StatusForbidden, r.Code)
		})
}
