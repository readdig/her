package handy

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"io"
)

type Token struct{}

func genToken() string {
	b := make([]byte, 16)
	_, _ = io.ReadFull(rand.Reader, b)
	val := hex.EncodeToString(b)
	return val
}

func genTokenHTML() template.HTML {
	name := "_xsrf"
	token := genToken()
	xsrfCookie := Config.Get("XSRFCookies").Bool()
	if token != "" && xsrfCookie {
		return template.HTML(fmt.Sprintf(`<input type="hidden" value="%s" name="%s" id="%s">`, name, token, name))
	}
	return template.HTML("")
}
