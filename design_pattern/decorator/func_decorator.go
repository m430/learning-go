package decorator

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

// 定义一个字符串操作器，类型为一个函数，可以传入一个string，同时返回的还是string
type StringOperator func(string) string

// Just仅仅返回字符串本身，不做处理
func Just(s string) string {
	fmt.Printf("Just: %s\n", s)
	return s
}

// 字符串变大写装饰器
func ToUpper(so StringOperator) StringOperator {
	return func(s string) string {
		str := strings.ToUpper(s)
		fmt.Printf("ToUpper: %s\n", str)
		return so(str)
	}
}

// MD5加密装饰器
func ToMD5(so StringOperator) StringOperator {
	return func(s string) string {
		h := md5.New()
		h.Write([]byte(s))
		md5 := hex.EncodeToString(h.Sum(nil))
		fmt.Printf("ToMD5: %s\n", md5)
		return so(md5)
	}
}

type Decorator func(StringOperator) StringOperator

// 定义一个pipline函数，可以按照次数依次进行装饰
func Pipline(so StringOperator, decorstors ...Decorator) StringOperator {
	for i := len(decorstors) - 1; i >= 0; i-- {
		so = decorstors[i](so)
	}
	return so
}
