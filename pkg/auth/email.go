package auth

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"xiaolong_blog/global"
)

const emailMessage = `
<div style="max-width:800px;letter-spacing: 0.2px;padding: 2rem;box-shadow: 0 0 10px #eee;">
    <h2 style="color: #303030;font-size: 1rem;font-weight: 400;">Hi 你好,我是小龙!</h2>
       <p style="font-size: .9rem;line-height: 24px;">感谢你订阅了我的博客
我会不定时的更新文章,希望你喜欢。</p>
                    <div style="margin:20px 20px;"> 验证码:<span style="font-size:24px;"><strong>%s</strong></span></div>
                    <div
                        style="background: #eff5fb;border-left: 4px solid #c2e1ff;padding: 14px;margin-top: 30px;border-radius: 9px;font-size: 0.85rem;color: #7d7f7f;line-height: 24px;">
                        If we don't have a chance to meet, then I'm here to wish you good morning, good afternoon and good night in advance～～<br>愿所有的美好如约而至，愿所有的黑暗都能看到希望，我们微笑前行～～<br>人生没有完美，也许有些遗憾才美～～永远相信美好的事情即将发生～～
                    </div>
                </div>
            `

func SendEmail(recipient string, verifyCode string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", global.EmailSetting.UserName)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", "邮箱验证")
	m.SetBody("text/html", fmt.Sprintf(emailMessage, verifyCode))

	d := gomail.NewDialer(global.EmailSetting.Host, global.EmailSetting.Port, global.EmailSetting.UserName, global.EmailSetting.Password)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false}

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return nil

}
