package her

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

func genTokenHTML(ctx *Context) template.HTML {
	name := "_xsrf"
	token := ctx.GetToken()
	xsrfCookie := Config.Bool("XSRFCookies")
	if token != "" && xsrfCookie {
		ctx.SetCookie(NewCookie(name, token, 24))
		return template.HTML(fmt.Sprintf(`<input type="hidden" value="%s" name="%s">`, token, name))
	}
	return template.HTML("")
}
