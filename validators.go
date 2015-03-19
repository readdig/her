package her

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const RE_RFC_EMAIL = `(?i)[A-Z0-9!#$%&'*+/=?^_{|}~-]+` +
	`(?:\.[A-Z0-9!#$%&'*+/=?^_{|}~-]+)*` +
	`@(?:[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?\.)+` +
	`[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?`

type Validator interface {
	CleanData(value string) (bool, string)
}

type Required struct {
	Message string
}

func (v Required) CleanData(value string) (bool, string) {
	if value == "" {
		message := "字段是必填的。"
		if v.Message != "" {
			message = v.Message
		}
		return false, message
	}

	return true, ""
}

type Regexp struct {
	Expr    string
	Message string
}

func (v Regexp) CleanData(value string) (bool, string) {
	reg, err := regexp.Compile(v.Expr)
	if err != nil {
		panic(err)
	}

	if reg.MatchString(value) {
		return true, ""
	}

	return false, v.Message
}

type Email struct {
	Message string
}

func (v Email) CleanData(value string) (bool, string) {
	message := "无效的 Email 地址。"
	if v.Message != "" {
		message = v.Message
	}
	tmp := Regexp{Expr: RE_RFC_EMAIL, Message: message}

	return tmp.CleanData(value)
}

type URL struct {
	Message string
}

func (v URL) CleanData(value string) (bool, string) {
	message := "无效的 URL。"
	if v.Message != "" {
		message = v.Message
	}
	tmp := Regexp{Expr: `^(http|https)?://([^/:]+|([0-9]{1,3}\.){3}[0-9]{1,3})(:[0-9]+)?(\/.*)?$`, Message: message}

	return tmp.CleanData(value)
}

type Length struct {
	Min     int
	Max     int
	Message string
}

func (v Length) CleanData(value string) (bool, string) {
	message := ""
	minMessage := "字段长度必须至少 %d 个字符。"
	maxMessage := "字段长度不能超过 %d 个字符。"
	minAndMaxMessage := "字段长度必须介于 %d 到 %d 个字符之间。"

	if v.Message != "" {
		message = v.Message
	}
	if len(value) < v.Min && v.Max == 0 {
		if message == "" {
			message = minMessage
		}
		return false, fmt.Sprintf(message, v.Min)
	}
	if len(value) > v.Max && v.Min == 0 {
		if message == "" {
			message = maxMessage
		}
		return false, fmt.Sprintf(message, v.Max)
	}
	if len(value) < v.Min || v.Max > 0 && len(value) > v.Max {
		if message == "" {
			message = minAndMaxMessage
		}
		return false, fmt.Sprintf(message, v.Min, v.Max)
	}
	return true, ""
}

type NumberRange struct {
	Max     int
	Min     int
	Message string
}

func (v NumberRange) CleanData(value string) (bool, string) {

	message := ""
	minMessage := "值必须大于 %d 。"
	maxMessage := "值必须小于 %d 。"
	minAndMaxMessage := "值必须介于 %d 到 %d 个之间。"

	if v.Message != "" {
		message = v.Message
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		message = "必须是数字。"
		return false, message
	}

	//	println(val < v.Min && v.Max == 0)
	if val < v.Min && v.Max == 0 {
		if message == "" {
			message = minMessage
		}
		return false, fmt.Sprintf(message, v.Min)
	}
	if val > v.Max && v.Min == 0 {
		if message == "" {
			message = maxMessage
		}
		return false, fmt.Sprintf(message, v.Max)
	}
	if val < v.Min || v.Max > 0 && val > v.Max {
		if message == "" {
			message = minAndMaxMessage
		}
		return false, fmt.Sprintf(message, v.Min, v.Max)
	}

	return true, ""
}

type IPAddress struct {
	Message string
}

func (v IPAddress) CleanData(value string) (bool, string) {
	message := "无效的 IP 地址。"
	if v.Message != "" {
		message = v.Message
	}
	ipArray := strings.Split(value, ".")
	length := len(ipArray)
	if length != 4 {
		return false, message
	}

	for i := 0; i < 4; i++ {
		temp, err := strconv.Atoi(ipArray[i])
		if err != nil {
			return false, message
		}
		if len(ipArray[i]) == 0 || temp > 255 {
			return false, message
		}
	}
	return true, ""
}
