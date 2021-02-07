package main

import (
	"flag"
	"fmt"
	"github.com/go-gomail/gomail"
	"strings"
)

type EmailParam struct {
	// ServerHost 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	ServerHost string
	// ServerPort 邮箱服务器端口，如腾讯企业邮箱为465
	ServerPort int
	// FromEmail　发件人邮箱地址
	FromEmail string
	// FromPasswd 发件人邮箱密码（注意，这里是明文形式），TODO：如果设置成密文？
	FromPasswd string
	// Toers 接收者邮件，如有多个，则以英文逗号(“,”)隔开，不能为空
	Toers string
	// CCers 抄送者邮件，如有多个，则以英文逗号(“,”)隔开，可以为空
	CCers string
}

// 全局变量，因为发件人账号、密码，需要在发送时才指定
// 注意，由于是小写，外面的包无法使用
var serverHost, fromEmail, fromPasswd string
var serverPort int

var m *gomail.Message

func InitEmail() {

	// 结构体赋值
	ep := &EmailParam {
		ServerHost: "smtp.163.com",
		ServerPort: 465,
		FromEmail:  "giouya@163.com",
		FromPasswd: "CMOJHASAPVSICYHZ",
		Toers:      "65489469@qq.com, giouya@163.com",
		CCers:      "",
	}


	toers := []string{}

	serverHost = ep.ServerHost
	serverPort = ep.ServerPort
	fromEmail = ep.FromEmail
	fromPasswd = ep.FromPasswd

	m = gomail.NewMessage()

	if len(ep.Toers) == 0 {
		return
	}

	for _, tmp := range strings.Split(ep.Toers, ",") {
		toers = append(toers, strings.TrimSpace(tmp))
	}

	// 收件人可以有多个，故用此方式
	m.SetHeader("To", toers...)

	//抄送列表
	if len(ep.CCers) != 0 {
		for _, tmp := range strings.Split(ep.CCers, ",") {
			toers = append(toers, strings.TrimSpace(tmp))
		}
		m.SetHeader("Cc", toers...)
	}

	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	m.SetAddressHeader("From", fromEmail, "服务器")
}

// SendEmail body支持html格式字符串
func SendEmail(subject, body string) {
	// 主题
	m.SetHeader("Subject", subject)

	// 正文
	m.SetBody("text/html", body)

	d := gomail.NewPlainDialer(serverHost, serverPort, fromEmail, fromPasswd)
	// 发送
	err := d.DialAndSend(m)
	if err != nil {
		fmt.Println("发送邮件失败：",err)
		panic(err)
	}
}

var subjectParam = flag.String("subject", "", "input your email subject")
var bodyParam = flag.String("body", "", "input your email body")

func main() {

	flag.Parse()

	subject := *subjectParam
	body := *bodyParam

	fmt.Printf("【subject=%s】【body=%s】\n",subject,body)

	if subject == "" || body == ""{
		fmt.Println("输入为空，发送邮件失败!")
		return
	}

	InitEmail()
	SendEmail(subject, body)
}
