package template

import "fmt"

// 生成一次性密码，发送给用户。
// 通过短信和邮件发送，发送的步骤是一样的。
// 所有，可以抽象出来具体的步骤。
// 共用相同的模版方法，从而减少冗余代码。

// OTP (One Time Password)
// 抽象接口：具体步骤
type OTP interface {
	random() string
	cache(str string)
	message(msg string) string
	send(msg string) error
	publishMetric()
}

type template struct {
	otp OTP
}

// 模板方法
func (t *template) doOTP() error {
	otp := t.otp.random()
	t.otp.cache(otp)
	message := t.otp.message(otp)
	err := t.otp.send(message)
	if err != nil {
		return err
	}
	t.otp.publishMetric()
	return nil
}

// SMS Concrete implementation
type sms struct {
	template
}

func (s *sms) random() string {
	str := "1231"
	fmt.Printf("SMS: generating one time password:%s\n", str)
	return str
}

func (s *sms) cache(str string) {
	fmt.Printf("SMS: saving one time password: %s to cache\n", str)
}

func (s *sms) message(str string) string {
	return "SMS one time password for login is " + str
}

func (s *sms) send(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

func (s *sms) publishMetric() {
	fmt.Printf("SMS: publishing metrics\n")
}

// EMAIL Concrete implementation
type email struct {
	template
}

func (e *email) random() string {
	str := "1231"
	fmt.Printf("EMAIL: generating one time password %s\n", str)
	return str
}

func (e *email) cache(str string) {
	fmt.Printf("EMAIL: saving one time password: %s to cache\n", str)
}

func (e *email) message(str string) string {
	return "EMAIL one time password for login is " + str
}

func (e *email) send(message string) error {
	fmt.Printf("EMAIL: sending sms: %s\n", message)
	return nil
}

func (e *email) publishMetric() {
	fmt.Printf("EMAIL: publishing metrics\n")
}
