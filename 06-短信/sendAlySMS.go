package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"time"
)

type BackupsMsg struct {
	Name string `json:"name"`
	Time string `json:"time"`
}

type AbnormalMsg struct {
	Name string `json:"name"`
}

var phoneParam = flag.String("phone", "", "input your sms phones")
var bodyParam = flag.String("body", "", "input your sms body")
var typeParam = flag.String("type", "", "input your send type")


func main() {

	flag.Parse()

	phone := *phoneParam
	body := *bodyParam
	sendType := *typeParam

	fmt.Printf("【sendType=%s】【phone=%s】【body=%s】\n",sendType,phone,body)

	if phone == "" || body == ""{
		fmt.Println("手机号或内容不能为空")
		return
	}

	if sendType == ""{
		fmt.Println("发送类型不能为空")
		return
	}

	if sendType == "backup"{
		//备份
		backupMsg := &BackupsMsg{
			Name: body,
			Time: time.Now().Format("2006-01-02 15:04:05"),
		}

		bys,_ := json.Marshal(backupMsg)
		sendBackupsSMS(phone,string(bys))


	}else if sendType == "abnormal"{
		//异常
		abnormalMsg := &AbnormalMsg{
			Name: body,
		}
		bys2,_ := json.Marshal(abnormalMsg)

		sendAbnormalSMS(phone,string(bys2))
	}else {
		fmt.Println("输入类型有误！")
		return
	}


}

func sendBackupsSMS(phone string,data string)  {
	client, err := dysmsapi.NewClientWithAccessKey(
		"cn-hangzhou",
		"LTAIh6K3QYEOtQMo",
		"56jYl2vckHmKnduwQt4mAjLPfYNf64")

	request := dysmsapi.CreateSendSmsRequest()

	request.PhoneNumbers = phone
	request.SignName = "哈珥斯"
	request.TemplateCode = "SMS_209836628"

	request.TemplateParam = data

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}

func sendAbnormalSMS(phone string,data string)  {
	client, err := dysmsapi.NewClientWithAccessKey(
		"cn-hangzhou",
		"LTAIh6K3QYEOtQMo",
		"56jYl2vckHmKnduwQt4mAjLPfYNf64")

	request := dysmsapi.CreateSendSmsRequest()

	request.PhoneNumbers = phone
	request.SignName = "哈珥斯"
	request.TemplateCode = "SMS_209836629"
	request.TemplateParam = data

	response, err := client.SendSms(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
}
