# 模板方法模式

模板方法是一种行为设计模式，它在超类中定义了一个算法的框架，允许子类在不修改结构的情况下重写算法的特定步骤。

Template Method is a behavioral design pattern that defines the skeleton(框架) of an algorithm in the superclass but lets
subclasses override specific steps of the algorithm without changing its structure.

![image](https://user-images.githubusercontent.com/65383410/165442902-28625ab1-c6e0-4639-8233-b3e82288947e.png)

Template method breaks the algorithm into steps, allowing subclasses to override these steps but not the actual method.

So, we have a scenario where the steps of a particular operation are the same, but these steps’ implementation may
differ. This is an appropriate situation to consider using the Template Method pattern.

We define a base template algorithm that consists of a fixed number of methods. That’ll be our template method. We will
then implement each of the step methods, but leave the template method unchanged.

## 实现

```go
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

```

## 用法

```go
package template

import (
	"fmt"
	"testing"
)

func TestTemplate(t *testing.T) {
	sms := template{
		otp: &sms{},
	}
	sms.doOTP()

	fmt.Println()

	email := template{
		otp: &email{},
	}
	email.doOTP()
}

Output:
SMS: generating one time password:1231
SMS: saving one time password: 1231 to cache
SMS: sending sms: SMS one time password for login is 1231
SMS: publishing metrics

EMAIL: generating one time password 1231
EMAIL: saving one time password: 1231 to cache
EMAIL: sending sms: EMAIL one time password for login is 1231
EMAIL: publishing metrics

```

## 应用场景

- 当你只希望客户端扩展某个特定算法步骤，而不是整个算法或其结构时，可使用模板方法模式。
    - 模板方法将整个算法转换为一系列独立的步骤， 以便子类能 对其进行扩展， 同时还可让超类中所定义的结构保持完整。

- 当多个类的算法除一些细微不同之外几乎完全一样时，你可 使用该模式。 但其后果就是， 只要算法发生变化， 你就可能 需要修改所有的类。
    - 在将算法转换为模板方法时， 你可将相似的实现步骤提取到 超类中以去除重复代码。 子类间各不同的代码可继续保留在 子类中。

## 优点

- 你可仅允许客户端重写一个大型算法中的特定部分， 使得算法其他部分修改对其所造成的影响减小。
- 你可将重复代码提取到一个超类中。

## 缺点

- 部分客户端可能会受到算法框架的限制。
- 模板方法中的步骤越多， 其维护工作就可能会越困难。
