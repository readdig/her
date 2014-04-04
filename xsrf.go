package her

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"html/template"
	"io"
)

const (
	tokenName          = "_xsrf"
	tokenCookieExpires = 1800
)

func genToken() string {
	b := make([]byte, 16)
	_, _ = io.ReadFull(rand.Reader, b)
	val := hex.EncodeToString(b)
	return val
}

func genTokenHTML(ctx *Context) template.HTML {
	token := ctx.GetToken()
	xsrfCookie := Config.Bool("XSRFCookies")
	if token != "" && xsrfCookie {
		ctx.SetCookie(tokenName, token, tokenCookieExpires)
		return template.HTML(fmt.Sprintf(`<input type="hidden" value="%s" name="%s">`, token, tokenName))
	}
	return template.HTML("")
}

func validateToken(ctx *Context) bool {
	token := ctx.GetToken()
	tokenXSRF := ctx.Request.FormValue(tokenName)
	if tokenXSRF == "" {
		return false
	}
	if token != tokenXSRF {
		return false
	}
	return true
}
