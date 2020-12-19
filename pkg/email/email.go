package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

type SMTPInfo struct {
	Host string
	Port int
	IsSSL bool
	UserName string
	Password string
	From string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo:info}
}

func (e *Email) SendMail(to []string, subject, body string) error {
	message := gomail.NewMessage()
	message.SetHeader("Form",e.From)
	message.SetHeader("To",to...)
	message.SetHeader("Subject",subject)
	message.SetBody("text/html",body)

	// 调用 NewDialer 方法创建一个新的SMTP 拨号实例，设置对应的拨号信息，用于连接SMTP服务器最后调用DialAndSend方法，打开余SMTP服务器的连接并发送电子邮件
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig =&tls.Config{
		InsecureSkipVerify:          e.IsSSL,
	}

	return dialer.DialAndSend(message)
}