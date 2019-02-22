package main

import (
	"fmt"
	intersms "submail_go_sdk/submail/internationalsms"
	mail "submail_go_sdk/submail/mail"
	sms "submail_go_sdk/submail/sms"
	voice "submail_go_sdk/submail/voice"
) 

func VoiceSend() {
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="you_appkey"
	config["signType"]="normal"

	submail := voice.CreateSend(config)
	submail.SetTo("138********")
	submail.SetContent("语音测试")
	send := submail.Send()
	fmt.Println("语音 Send 接口:",send)
}

func InterSMSSend(){
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="you_appkey"
	config["signType"]="normal"

	submail := intersms.CreateSend(config)

	//设置联系人 手机号码
	submail.SetTo("+1xxxxxxx")
	
	submail.SetContent("【SUBMAIL】test")

	//执行 Send 方法发送短信
	send := submail.Send()
	fmt.Println("国际短信 Send 接口:",send)

}

func MailSend()  {
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="you_appkey"
	config["signType"]="normal"

	submail := mail.CreateSend(config)
	submail.SetForm("service@submail.cn","submail")
	submail.AddTo("xxx@qq.com","Leo")
	submail.SetSubject("test from go sdks 2")
	submail.AddCc("xxx@gmail.com")
	submail.AddBcc("xxx@me.com")
	submail.SetText("text from go sdks")
	submail.AddVar("name","leo")
	submail.AddLink("url","https://www.mysubmail.com")
	submail.AddHeaders("X-Mailer","SUBMAIL Golang SDK")
	submail.AddAttachments("/Users/leozhang/go/src/submail_go_sdk/test_attachment.png")
	submail.AddAttachments("/Users/leozhang/go/src/submail_go_sdk/test_attachment2.jpg")
	send := submail.Send()
	fmt.Println("邮件 Send 接口:",send)
}

func SMSSend() {
	// SMS 短信服务配置 appid & appkey 请前往：https://www.mysubmail.com/chs/sms/apps 获取
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="you_appkey"
	// SMS 数字签名模式 normal or md5 or sha1 ,normal = 明文appkey鉴权 ，md5 和 sha1 为数字签名鉴权模式
	config["signType"]="normal"

	//创建 短信 Send 接口
	submail := sms.CreateSend(config)

	//设置联系人 手机号码
	submail.SetTo("138********")

	//设置短信正文，请注意：国内短信需要强制添加短信签名，并且需要使用全角大括号 “【签名】”标识，并放在短信正文的最前面
	submail.SetContent("【SUBMAIL】您的验证码是：2234，请在30分钟输入")

	//执行 Send 方法发送短信
	send := submail.Send()
	fmt.Println("短信 Send 接口:",send)
}

func main()  {
	//SMSSend()
	//MailSend()
	//VoiceSend()
	//InterSMSSend()
}