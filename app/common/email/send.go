package email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"log"
	"miniproject/app/common/tool"
	"miniproject/app/model"
	"miniproject/app/service/signup"
	"net/smtp"
)

func Send(usermail string) {
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "muxi < 发送者@qq.com>"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{usermail}

	// 设置主题
	em.Subject = "test subject"
	code, err := tool.GenerateRandomString(6)
	if err != nil {
		fmt.Println("err(generate code):", err)
	}

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.Text = code
	//将验证码和email都存入redis中
	code1 := string(code)
	signup.Savecode(model.Client, usermail, code1)
	model.Code = signup.Getcode(model.Client, usermail)
	//设置服务器相关的配置
	err = em.Send("smtp.qq.com:25", smtp.PlainAuth("", "发送者@qq.com", "授权码", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("send successfully ... ")
}
