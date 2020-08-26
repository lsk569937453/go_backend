package mail

import (
	"gopkg.in/gomail.v2"
)

var (
	mailTo = []string{ // 收件人列表
		`569937453@qq.com`,

	}
	title = `测试邮件标题` // 邮件主题 或者 邮件标题
	body  = `测试邮件内容` // 邮件内容（支持HTML）

	/* ====== 分割线 ====== */

	user = `123-abc@qq.com` // 发送邮箱：账号
	pass = `xxxxxx`            // 发送邮箱：密码（qq邮箱：密码填授权码）
	host = `smtp.qq.com`       // 发送邮箱：服务器地址
	port = 25                  // 发送邮箱：端口（默认端口：465，QQ邮箱端口：25）
)
/**

 */
func main() {
	m := gomail.NewMessage()
	m.SetHeader(`From`, user)
	m.SetHeader(`To`, mailTo...)
	m.SetHeader(`Subject`, title)
	m.SetBody(`text/html`, body)
	err := gomail.NewDialer(host, port, user, pass).DialAndSend(m)
	if err != nil {
		panic(err)
	}
}
