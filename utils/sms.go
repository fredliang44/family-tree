package utils

import (
	"fmt"
	"github.com/GiterLab/aliyun-sms-go-sdk/dysms"
	"github.com/gin-gonic/gin"
	"github.com/qichengzx/qcloudsms_go"
	"github.com/tobyzxj/uuid"
)

// SendDYSMS is a func to handle sms with dayu
func SendDYSMS(Phone, Template, Dict string) (isOK bool, msg string, errID string) {
	if gin.Mode() == "test" {
		if Phone == "17777766667" || Phone == "" {
			return false, "Testing err sms", ""
		}
		return true, "Done", ""
	}
	dysms.HTTPDebugEnable = true
	dysms.SetACLClient(AppConfig.Dayu.AccessID, AppConfig.Dayu.AccessKey)

	// send to some one
	respSendSms, err := dysms.SendSms(uuid.New(), Phone, AppConfig.Dayu.Sign, Template, Dict).DoActionWithException()
	respond := respSendSms.GetRequestID()
	if (err != nil) || (respSendSms.GetMessage() != "OK") {
		fmt.Println("send sms failed", err, respSendSms.Error())
		return false, respSendSms.GetMessage(), respond
	}
	return true, "Done", respond
}

// SendDYSMS is a func to handle sms with qcloud
func SendQCSMS(Phone string, Template int, ParamList []string) (isOK bool, msg string, errID string) {
	if gin.Mode() == "test" {
		if Phone == "17777766667" || Phone == "" {
			return false, "Testing err sms", ""
		}
		return true, "Done", ""
	}
	opt := qcloudsms.NewOptions(AppConfig.QcloudSMS.AppID, AppConfig.QcloudSMS.AppKey, AppConfig.QcloudSMS.Sign)

	var client = qcloudsms.NewClient(opt)
	client.SetDebug(true)

	var t = qcloudsms.SMSSingleReq{
		Params: ParamList,
		Tel:    qcloudsms.SMSTel{Nationcode: "86", Mobile: Phone},
		Sign:   AppConfig.QcloudSMS.Sign,
		TplID:  Template,
	}

	isOK, err := client.SendSMSSingle(t)
	return isOK, fmt.Sprintln(err), fmt.Sprintln(err)
}
