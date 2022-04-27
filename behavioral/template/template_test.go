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
