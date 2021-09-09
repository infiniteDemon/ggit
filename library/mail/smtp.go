package mail

import (
	"github.com/go-gomail/gomail"
	"service-all/library/logger"
)

type Econf struct {
	// 发件方配置
	Email    string
	Password string
	Host     string
	Port     int
	// 收件方配置
	From    string
	To      string
	Subject string
	Body    string
}

type Email interface {
	SendMail() error
}

func (e Econf) SendMail() error {
	// 发送邮件
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", e.To)
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", e.Subject)
	m.SetBody("text/html", e.Body)
	//m.Attach("/home/Alex/lolcat.jpg")

	logger.Log().Info("%v %v %v %v", e.Host, e.Port, e.Email)
	d := gomail.NewDialer(e.Host, e.Port, e.Email, e.Password)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		logger.Log().Info("邮件发送失败 %s", err)
		return err
	}
	return nil
}
