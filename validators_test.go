package handy

import (
	"testing"
)

type emailValidateTest struct {
	email string
	ok    bool
}

var emailvalidatetests = []emailValidateTest{
	{"username@example.com", true},
	{"aa.bb@example.com", true},
	{"aa_bb@example.com", true},
	{"aa-bb@example.com", true},
	{"aa123@example.com", true},
	{"AabB@example.com", true},
	{"Aa123.bb123@example.edu.cn", true},
	{"a@example.com", true},
	{"aa@example", false},
	{"<script>alert(1);</script>@test.com", false},
	{"aabbcc", false},
	{"aa@.com", false},
	{"aa.@xx.com", false},
	{"aa@", false},
	{"aa@bb@.example.com", false},
}

func TestEmailValidator(t *testing.T) {
	var email = Email{}

	for _, test := range emailvalidatetests {
		if ok, _ := email.CleanData(test.email); ok != test.ok {
			t.Error("test Email.CleanData:", test.email)
		}
	}
}

func TestIPAddress(t *testing.T) {

	ip := IPAddress{}

	ok, err := ip.CleanData("5.2.3.1")
	if err != "" {
		println(err)
		return
	}
	println("IPAddres Result:", ok)

}

func TestNumberRange(t *testing.T) {
	num := NumberRange{Min: 10, Max: 15}

	ok, err := num.CleanData(10)
	if err != "" {
		println("Error", err)
		return
	}
	println("NumberRange Result:", ok)

}
