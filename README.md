# **GO语言SDK**
#### 引入方式：
1. 下载SDK源码，然后放在$GOPATH/src目录下.
1. 修改文件夹名为submail_go_sdk。
1. 在项目下进行包引用。如短信包：import (sms "submail_go_sdk/submail/sms")  / 邮件包 import (mail "submail_go_sdk/submail/mail")
---


#### 文件目录索引 
##### Submail:
	internationalsms
	    multixsend.go          国际短信群发接口
	    send.go                国际短信send接口
	    xsend.go               国际短信xsend接口

     mail
	    send.go                邮件send接口
	    xsend.go               邮件xsend接口

	 sms
	    multixsend.go          国内短信群发接口
	    send.go                国内短信send接口
	    xsend.go               国内短信xsend接口
		
	 voice
	    multixsend.go          语音群发接口
	    send.go                语音send接口
	    xsend.go               语音xsend接口
	    
	 lib
	    lib.go                 处理http请求以及生成signature参数

---

## 开始使用 
###邮件 Mail 包：
    邮件服务 appid & appkey 请前往：https://www.mysubmail.com/chs/mail/apps
  
    mail 数字签名模式 normal or md5 or sha1 ,normal = 明文appkey鉴权 ，md5 和 sha1 为数字签名鉴权模式

		
### 邮件 Send 接口:
##### 服务配置: 
- config := make(map[string]string)
- config["appid"]="your_appid"
- config["appkey"]="your_appkey"
- config["signType"]="sha1"

 

**使用指引:**

方法名 | 描述 
---|---
CreateSend | 创建邮件Send 接口
AddTo | 添加邮件地址到 To map，第一个必选参数：邮件地址。第二个可选参数：收件人姓名
SetSender | 设置发件人，第一个必选参数：邮件地址。第二个可选参数：显示名称
SetReply | 设置回复地址
AddCc | 添加抄送地址
AddBcc | 添加密送地址
SetSubject | 设置邮件标题
SetText | 设置文本邮件内容
SetHtml | 设置 HTML 邮件内容
AddVar | 添加文本变量到 vars map
AddLink | 添加超链接变量到 links map
AddHeaders | 添加自定义邮件头指令到 headers map
AddAttachments | 添加附件到 attachments map，传多个附件多次调用即可。
SetAsynchronous | 设置异步选项，该值设为 true 时启用异步发送模式
SetTag |  设置自定义标记，此参数用于标记一次 API 请求（最大长度不超过 32 位）添加了 tag 参数的 API 请求，会在所有的 SUBHOOK 事件中携带此参数。
Send | 发送邮件

**代码示列：**

```
使用 MailSend 类提交 mail/send 发送一封简单的邮件
func MailSend()  {
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="your_appkey"
	config["signType"]="sha1"
	submail := mail.CreateSend(config)
	submail.SetSender("service@submail.cn","submail")
	submail.AddTo("79463@qq.com","Leo")
	submail.SetSubject("test from go sdks 2")
	submail.AddCc("drimmestudio@gmail.com")
	submail.AddBcc("drimmestudio@me.com")
	submail.SetText("text from go sdks")
	submail.AddVar("name","leo")
	submail.AddLink("url","https://www.mysubmail.com")
	submail.AddHeaders("X-Mailer","SUBMAIL Golang SDK")
	submail.AddAttachments("/Users/leozhang/go/src/submail_go_sdk/test_attachment.png")
	submail.AddAttachments("/Users/leozhang/go/src/submail_go_sdk/test_attachment2.jpg")
	send := submail.Send()
	fmt.Println("Request result:",send)
}

```

---

### 邮件 xsend 接口:
- config := make(map[string]string)
- config["appid"]="your_appid"
- config["appkey"]="your_appkey"
- config["signType"]="sha1"

**使用指引:**

方法  | 描述 
---|---
CreateXsend | 创建邮件xsend 接口
AddTo | 添加邮件地址到 To map，第一个必选参数：邮件地址。第二个可选参数：收件人姓名
SetSender | 设置发件人，第一个必选参数：邮件地址。第二个可选参数：显示名称
SetReply | 设置回复地址
AddCc | 添加抄送地址
AddBcc | 添加密送地址
SetSubject | 设置邮件标题
SetProject | 设置邮件项目标识
AddVar | 添加文本变量到 vars map
AddLink | 添加超链接变量到 links map
AddHeaders | 添加自定义邮件头指令到 headers map
SetAsynchronous | 设置异步选项，该值设为 true 时启用异步发送模式
SetTag |  设置自定义标记，此参数用于标记一次 API 请求（最大长度不超过 32 位）添加了 tag 参数的 API 请求，会在所有的 SUBHOOK 事件中携带此参数。
XSend | 发送邮件
**代码示例:**

```
使用 MAILXsend 类提交 mail/xsend 发送一封邮件。

func MailXsend()  {
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="your_appkey"
	config["signType"]="sha1"
	submail := mail.CreateXsend(config)
	submail.SetSender("service@submail.cn","submail")
	submail.AddTo("79463@qq.com","Leo")
	submail.SetSubject("test from go sdks 2")
	submail.SetProject("Zx3LM")
	submail.AddCc("drimmestudio@gmail.com")
	submail.AddBcc("drimmestudio@me.com")
	submail.AddVar("name","leo")
	submail.AddLink("url","https://www.mysubmail.com")
	submail.AddHeaders("X-Mailer","SUBMAIL Golang SDK")
	submail.SetTag("xxxxxx")
	submail.SetAsynchronous(true)
	xsend := submail.Xsend()
	fmt.Println("Request result:",xsend)
}

```

---
###短信 SMS 包
    短信服务 appid & appkey 请前往：https://www.mysubmail.com/chs/sms/apps
  
    短信数字签名模式 normal or md5 or sha1 ,normal = 明文appkey鉴权 ，md5 和 sha1 为数字签名鉴权模式


### 短信 send 接口:
##### 服务配置: 
- config := make(map[string]string)
- config["appid"]="your_appid"
- config["appkey"]="your_appkey"
- config["signType"]="sha1"

**使用指引:**

方法  | 描述 
---|---
CreateSend | 创建短信send接口
SetTo | 设置手机联系人
SetContent | 设置短信正文，请注意：国内短信需要强制添加短信签名，并且需要使用全角大括号 “【签名】”标识，并放在短信正文的最前面
SetTag |  设置自定义标记，此参数用于标记一次 API 请求（最大长度不超过 32 位）添加了 tag 参数的 API 请求，会在所有的 SUBHOOK 事件中携带此参数。
Send | 发送短信

**代码示列：**

```
  func SMSSend() {
	// SMS 短信服务配置 appid & appkey 请前往：https://www.mysubmail.com/chs/sms/apps 获取
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="your_appkey"
	// SMS 数字签名模式 normal or md5 or sha1 ,normal = 明文appkey鉴权 ，md5 和 sha1 为数字签名鉴权模式
	config["signType"]="sha1"
	//创建 短信 Send 接口
	submail := sms.CreateSend(config)
	//设置联系人 手机号码
	submail.SetTo("your_telephone")
	//设置短信正文，请注意：国内短信需要强制添加短信签名，并且需要使用全角大括号 “【签名】”标识，并放在短信正文的最前面
	submail.SetContent("【SUBMAIL】您的验证码是：2234，请在30分钟输入")
	//执行 Send 方法发送短信
	send := submail.Send()
	fmt.Println("短信 Send 接口:",send)
}

```

---

### 短信 Xsend 接口:
##### 服务配置: 
- config := make(map[string]string)
- config["appid"]="your_appid"
- config["appkey"]="your_appkey"
- config["signType"]="sha1"

**使用指引:**

方法  | 描述 
---|---
CreateXsend | 创建短信xsend接口
SetTo | 设置手机联系人
SetProject | 设置项目Id
AddVar | 添加文本变量到 vars map
SetTag | 设置自定义标记，此参数用于标记一次 API 请求（最大长度不超过 32 位）添加了 tag 参数的 API 请求，会在所有的 SUBHOOK 事件中携带此参数。
XSend | 发送短信

**代码示列：**

```
func SMSXsend() {
	// SMS 短信服务配置 appid & appkey 请前往：https://www.mysubmail.com/chs/sms/apps 获取
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="your_appkey"
	// SMS 数字签名模式 normal or md5 or sha1 ,normal = 明文appkey鉴权 ，md5 和 sha1 为数字签名鉴权模式
	config["signType"]="sha1"
	//创建 短信 Send 接口
	submail := sms.CreateXsend(config)
	//设置联系人 手机号码
	submail.SetTo("your_telephone")
    //设置短信模板id
	submail.SetProject("XV6HJ")
	//添加模板中的设置的动态变量。如模板为：【xxx】您的验证码是:@var(code),请在@var(time)分钟内输入。
	submail.AddVar("code","1234");
	submail.AddVar("time","5");
	//执行 Xsend 方法发送短信
	xsend := submail.Xsend()
	fmt.Println("短信XSend 接口:",xsend)

```

---

### 短信 mutilxsend 接口:
##### 服务配置: 
- config := make(map[string]string)
- config["appid"]="your_appid"
- config["appkey"]="your_appkey"
- config["signType"]="sha1"

**使用指引:**

方法  | 描述 
---|---
CreateMultiXsend | 创建短信CreateMultiXsend接口
SetProject | 设置项目Id
CreateMulti| 设置多个发送对象信息
SetTo | 设置联系人手机号码
AddVar | 添加模板变量
AddTag | 添加项目标记，此参数用于标记一次 API 请求（最大长度不超过 32 位）添加了 tag 参数的 API 请求，会在所有的 SUBHOOK 事件中携带此参数。
MultiXsend| 发送短信

**代码示列：**

```
func SMSMultiXsend() {
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="your_appkey"
	config["signType"]="sha1"
	submail := sms.CreateMultiXsend(config)
	submail.SetProject("XV6HJ")
	//添加第一组收件人
	multi1:=sms.CreateMulti();
	multi1.SetTo("133XXXXX");
	multi1.AddVar("code","1234");
	multi2.AddVar("time","10");
	//添加第二组收件人
	multi2:=sms.CreateMulti();
	multi2.SetTo("133XXXXX");
	multi2.AddVar("code","1234");
	multi2.AddVar("time","10");
	submail.AddMulti(multi1.Get());
	submail.AddMulti(multi2.Get());
	multixsend := submail.MultiXsend()
	fmt.Println("短信MultiXsend 接口:",multixsend)
```

---

##语音 Voice 包
    语音服务 appid & appkey 请前往：https://www.mysubmail.com/chs/voice/apps
  
    语音数字签名模式 normal or md5 or sha1 ,normal = 明文appkey鉴权 ，md5 和 sha1 为数字签名鉴权模式


### 语音 send 接口:
##### 服务配置: 
- config := make(map[string]string)
- config["appid"]="your_appid"
- config["appkey"]="your_appkey"
- config["signType"]="sha1"

**使用指引:**

方法  | 描述 
---|---
CreateSend | 创建短信send接口
SetTo | 设置手机联系人
SetContent | 设置语音正文
Send() | 发送短信

**代码示列：**

```
  func VoiceSend() {
	// 语音 短信服务配置 appid & appkey 请前往：https://www.mysubmail.com/chs/voice/apps 获取
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="your_appkey"
	// 语音数字签名模式 normal or md5 or sha1 ,normal = 明文appkey鉴权 ，md5 和 sha1 为数字签名鉴权模式
	config["signType"]="sha1"
	//创建语音Send 接口
	submail := sms.CreateSend(config)
	//设置联系人手机号码
	submail.SetTo("your_telephone")
	//设置语音正文。
	submail.SetContent("您的验证码是：2234，请在30分钟输入")
	//执行 Send 方法发送语音
	send := submail.Send()
	fmt.Printf("语音 Send 接口:%s\n",send)
}

```


### 语音 xsend 接口:
##### 服务配置: 
- config := make(map[string]string)
- config["appid"]="your_appid"
- config["appkey"]="your_appkey"
- config["signType"]="sha1"

**使用指引:**

方法  | 描述 
---|---
CreateXsend | 创建短信xsend接口
SetTo | 设置手机联系人
SetProject | 设置项目Id
AddVar | 添加文本变量到 vars map
XSend | 发送语音

**代码示列：**

```
func VoiceXsend() {
	// 语音服务配置 appid & appkey 请前往：https://www.mysubmail.com/chs/voice/apps 获取
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="your_appkey"
	// 语音 数字签名模式 normal or md5 or sha1 ,normal = 明文appkey鉴权 ，md5 和 sha1 为数字签名鉴权模式
	config["signType"]="sha1"
	//创建 语音 XSend 接口
	submail := sms.CreateXsend(config)
	//设置联系人 手机号码
	submail.SetTo("your_telephone")
    //设置语音模板id
	submail.SetProject("XV6HJ")
	//添加模板中的设置的动态变量。如模板为：您的验证码是:@var(code),请在@var(time)分钟内输入。
	submail.AddVar("code","1234");
	submail.AddVar("time","5");
	//执行 Xsend 方法发送语音
	xsend := submail.Xsend()
	fmt.Println("语音XSend 接口:",xsend)

```

---

### 语音 multixsend 接口:
##### 服务配置: 
- config := make(map[string]string)
- config["appid"]="your_appid"
- config["appkey"]="your_appkey"
- config["signType"]="sha1"

**使用指引:**

方法  | 描述 
---|---
CreateMultiXsend | 创建短信CreateMultiXsend接口
SetProject | 设置项目Id
CreateMulti| 设置多个发送对象信息
SetTo | 设置联系人手机号码
AddVar | 添加模板变量
MultiXsend| 发送语音

**代码示列：**

```
func VoiceMultiXsend() {
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="your_appkey"
	config["signType"]="sha1"
	submail := sms.CreateMultiXsend(config)
	submail.SetProject("XV6HJ")
	//添加第一组收件人
	multi1:=sms.CreateMulti();
	multi1.SetTo("133XXXXX");
	multi1.AddVar("code","1234");
	multi2.AddVar("time","10");
	//添加第二组收件人
	multi2:=sms.CreateMulti();
	multi2.SetTo("133XXXXX");
	multi2.AddVar("code","1234");
	multi2.AddVar("time","10");
	submail.AddMulti(multi1.Get());
	submail.AddMulti(multi2.Get());
	multixsend := submail.MultiXsend()
	fmt.Println("语音 MultiXsend 接口:",multixsend)
```

---

## 国际短信 internationalsms包
    国际短信服务 appid & appkey 请前往：https://www.mysubmail.com/chs/internationalsms/apps
  
    国际短信数字签名模式 normal or md5 or sha1 ,normal = 明文appkey鉴权 ，md5 和 sha1 为数字签名鉴权模式


### 国际短信 send 接口:
##### 服务配置: 
- config := make(map[string]string)
- config["appid"]="your_appid"
- config["appkey"]="your_appkey"
- config["signType"]="sha1"

**使用指引:**

方法  | 描述 
---|---
CreateSend | 创建国际短信send接口
SetTo | 设置手机联系人
SetContent | 设置短信正文。可以传自定义内容，非必须短信签名，纯英文短信（包括标点符号，短信签名的[]也需英文）单条按140个字符计费，超过140个字符每132个字符计费一次，其他语言单条按70个字符计费，超过70个字符每67个字符计费一次。
SetTag |  设置自定义标记，此参数用于标记一次 API 请求（最大长度不超过 32 位）添加了 tag 参数的 API 请求，会在所有的 SUBHOOK 事件中携带此参数。
Send | 发送短信

**代码示列：**

```
  func InternationalsmsSend() {
	// 国际短信服务配置 appid & appkey 请前往：https://www.mysubmail.com/chs/internationalsms/apps 获取
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="your_appkey"
	// 国际短信 数字签名模式 normal or md5 or sha1 ,normal = 明文appkey鉴权 ，md5 和 sha1 为数字签名鉴权模式
	config["signType"]="sha1"
	//创建 国际短信 Send 接口
	submail := sms.CreateSend(config)
	//设置联系人 手机号码
	submail.SetTo("your_telephone")
	//设置国际短信正文
	submail.SetContent("【SUBMAIL】您的验证码是：2234，请在30分钟输入")
	//执行 Send 方法发送国际短信
	send := submail.Send()
	fmt.Println("国际短信短信 Send 接口:",send)
}

```

---

### 国际短信 xsend 接口:
##### 服务配置: 
- config := make(map[string]string)
- config["appid"]="your_appid"
- config["appkey"]="your_appkey"
- config["signType"]="sha1"

**使用指引:**

方法  | 描述 
---|---
CreateXsend | 创建短信xsend接口
SetTo | 设置手机联系人
SetProject | 设置项目Id
AddVar | 添加文本变量到 vars map
SetTag | 设置自定义标记，此参数用于标记一次 API 请求（最大长度不超过 32 位）添加了 tag 参数的 API 请求，会在所有的 SUBHOOK 事件中携带此参数。
XSend | 发送短信

**代码示列：**

```
func  InternationalsmXsend() {
	// 国际短信服务配置 appid & appkey 请前往：https://www.mysubmail.com/chs/internationalsms/apps 获取
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="your_appkey"
	// 国际短信数字签名模式 normal or md5 or sha1 ,normal = 明文appkey鉴权 ，md5 和 sha1 为数字签名鉴权模式
	config["signType"]="sha1"
	//创建 国际短信 XSend 接口
	submail := sms.CreateXsend(config)
	//设置联系人 手机号码
	submail.SetTo("your_telephone")
    //设置国际短信模板id
	submail.SetProject("XV6HJ")
	//添加模板中的设置的动态变量。如模板为：【xxx】您的验证码是:@var(code),请在@var(time)分钟内输入。
	submail.AddVar("code","1234");
	submail.AddVar("time","5")
	//执行 Xsend 方法发送国际短信
	xsend := submail.Xsend()
	fmt.Println("国际短信 XSend 接口:",xsend)

```

---

### 国际短信 multixsend 接口:
##### 服务配置: 
- config := make(map[string]string)
- config["appid"]="your_appid"
- config["appkey"]="your_appkey"
- config["signType"]="sha1"

**使用指引:**

方法  | 描述 
---|---
CreateMultiXsend | 创建国际短信CreateMultiXsend接口
SetProject | 设置项目Id
CreateMulti| 设置多个发送对象信息
SetTo | 设置联系人手机号码
AddVar | 添加模板变量
AddTag | 添加项目标记，此参数用于标记一次 API 请求（最大长度不超过 32 位）添加了 tag 参数的 API 请求，会在所有的 SUBHOOK 事件中携带此参数。
MultiXsend| 发送短信

**代码示列：**

```
func Internationalsms MultiXsend() {
	config := make(map[string]string)
	config["appid"]="your_appid"
	config["appkey"]="your_appkey"
	config["signType"]="sha1"
	submail := sms.CreateMultiXsend(config)
	submail.SetProject("XV6HJ")
	//添加第一组收件人
	multi1:=sms.CreateMulti();
	multi1.SetTo("133XXXXX");
	multi1.AddVar("code","1234");
	multi2.AddVar("time","10");
	//添加第二组收件人
	multi2:=sms.CreateMulti();
	multi2.SetTo("133XXXXX");
	multi2.AddVar("code","1234");
	multi2.AddVar("time","10");
	submail.AddMulti(multi1.Get());
	submail.AddMulti(multi2.Get());
	multixsend := submail.MultiXsend()
	fmt.Println("国际短信 MultiXsend 接口:",multixsend)
```


具体 API 参数说明请参考 SUBMAIL 开发文档[：https://www.mysubmail.com/chs/documents/developer/index]()

