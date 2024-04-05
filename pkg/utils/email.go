package utils

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"sylu-oj-gin/internal/app/config"
	"sylu-oj-gin/pkg/logger"
)

func SendEmail(to string, code string) bool {
	// 构造邮件内容
	e := &email.Email{
		From:    fmt.Sprintf("%v <%v>", config.CONFIG.Email.Name, config.CONFIG.Email.Addr),
		To:      []string{to},
		Subject: "邮箱验证",
		Text:    []byte("注册验证码：" + code + "，有效期 10 分钟"),
	}

	err := e.Send(config.CONFIG.Email.Smtp, smtp.PlainAuth("", config.CONFIG.Email.Addr, config.CONFIG.Email.Password, config.CONFIG.Email.Host))
	if err != nil {
		logger.Error("验证码发送失败！", err)
		return false
	}

	return true
}
