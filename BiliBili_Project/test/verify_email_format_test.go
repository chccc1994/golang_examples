package test

import (
	"fmt"
	"testing"

	"github.com/chccc1994/bilibili/pkg"
)

func TestVerifyEmailFormat(t *testing.T) {
	fmt.Println("邮箱验证：", pkg.VerifyEmailFormat("12345@126.com")) //true
	fmt.Println("邮箱验证：", pkg.VerifyEmailFormat("12345126.com"))  //false
}
