package her

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Context struct {
	Request *http.Request
	http.ResponseWriter
	Params map[string]string
	Token  string
}

// WriteString writes string data into the response object.
func (ctx *Context) WriteString(content string) {
	ctx.Write([]byte(content))
}
func (ctx *Context) Status(status int) {
	ctx.WriteHeader(status)
}

// Abort is a helper method that sends an HTTP header and an optional
// body. It is useful for returning 4xx or 5xx errors.
// Once it has been called, any return value from the handler will
// not be written to the response.
func (ctx *Context) Abort(status int, body string) {
	ctx.WriteHeader(status)
	ctx.Write([]byte(body))
}

// Redirect is a helper method.
func (ctx *Context) Redirect(url string) {
	ctx.Header().Set("Location", url)
	ctx.WriteHeader(http.StatusFound)
	ctx.Write([]byte("Redirecting to: " + url))
}

func (ctx *Context) RedirectPermanent(url string) {
	ctx.Header().Set("Location", url)
	ctx.WriteHeader(http.StatusMovedPermanently)
	ctx.Write([]byte("Redirecting to: " + url))
}

// Notmodified writes a 304 HTTP response
func (ctx *Context) NotModified() {
	ctx.WriteHeader(304)
}

// NotFound writes a 404 HTTP response
func (ctx *Context) NotFound(message string) {
	ctx.WriteHeader(404)
	ctx.Write([]byte(message))
}

//Unauthorized writes a 401 HTTP response
func (ctx *Context) Unauthorized() {
	ctx.WriteHeader(401)
}

//Forbidden writes a 403 HTTP response
func (ctx *Context) Forbidden() {
	ctx.WriteHeader(403)
}

// ContentType sets the Content-Type header for an HTTP response.
// For example, ctx.ContentType("json") sets the content-type to "application/json"
// If the supplied value contains a slash (/) it is set as the Content-Type
// verbatim. The return value is the content type as it was
// set, or an empty string if none was found.
func (ctx *Context) ContentType(val string) string {
	var ctype string
	if strings.ContainsRune(val, '/') {
		ctype = val
	} else {
		if !strings.HasPrefix(val, ".") {
			val = "." + val
		}
		ctype = mime.TypeByExtension(val)
	}
	if ctype != "" {
		ctx.Header().Set("Content-Type", ctype)
	}
	return ctype
}

// SetHeader sets a response header. If `unique` is true, the current value
// of that header will be overwritten . If false, it will be appended.
func (ctx *Context) SetHeader(hdr string, val string, unique bool) {
	if unique {
		ctx.Header().Set(hdr, val)
	} else {
		ctx.Header().Add(hdr, val)
	}
}

// token is xsrf
func (ctx *Context) GetToken() string {
	return ctx.Token
}

// SetCookie adds a cookie header to the response.
func (ctx *Context) SetCookie(cookie *http.Cookie) {
	ctx.SetHeader("Set-Cookie", cookie.String(), false)
}

// NewCookie is a helper method that returns a new http.Cookie object.
// Duration is specified in seconds. If the duration is zero, the cookie is permanent.
// This can be used in conjunction with ctx.SetCookie.
func NewCookie(name string, value string, age int64) *http.Cookie {
	var utctime time.Time
	if age == 0 {
		// 2^31 - 1 seconds (roughly 2038)
		utctime = time.Unix(2147483647, 0)
	} else {
		utctime = time.Unix(time.Now().Unix()+age, 0)
	}
	return &http.Cookie{Name: name, Value: value, Expires: utctime}
}

func getCookieSig(key string, val []byte, timestamp string) string {
	hm := hmac.New(sha1.New, []byte(key))

	hm.Write(val)
	hm.Write([]byte(timestamp))

	hex := fmt.Sprintf("%02x", hm.Sum(nil))
	return hex
}

func (ctx *Context) SetSecureCookie(name string, val string, age int64) {
	cookieSecret := Config.String("CookieSecret")
	//base64 encode the val
	if len(cookieSecret) == 0 {
		return
	}
	var buf bytes.Buffer
	encoder := base64.NewEncoder(base64.StdEncoding, &buf)
	encoder.Write([]byte(val))
	encoder.Close()
	vs := buf.String()
	vb := buf.Bytes()
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	sig := getCookieSig(cookieSecret, vb, timestamp)
	cookie := strings.Join([]string{vs, timestamp, sig}, "|")
	ctx.SetCookie(NewCookie(name, cookie, age))
}

func (ctx *Context) GetSecureCookie(name string) (string, bool) {
	for _, cookie := range ctx.Request.Cookies() {
		if cookie.Name != name {
			continue
		}

		parts := strings.SplitN(cookie.Value, "|", 3)

		val := parts[0]
		timestamp := parts[1]
		sig := parts[2]

		cookieSecret := Config.String("CookieSecret")

		if getCookieSig(cookieSecret, []byte(val), timestamp) != sig {
			return "", false
		}

		ts, _ := strconv.ParseInt(timestamp, 0, 64)

		if time.Now().Unix()-31*86400 > ts {
			return "", false
		}

		buf := bytes.NewBufferString(val)
		encoder := base64.NewDecoder(base64.StdEncoding, buf)

		res, _ := ioutil.ReadAll(encoder)
		return string(res), true
	}
	return "", false
}

func (ctx *Context) Render(tmpl string, data map[string]interface{}) {
	if tmpl != "" {
		data["xsrf_form_html"] = genTokenHTML(ctx)
		err := templates.ExecuteTemplate(ctx.ResponseWriter, tmpl, data)
		if err != nil {
			http.Error(ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (ctx *Context) Json(v interface{}) {
	content, err := json.Marshal(v)
	if err == nil {
		ctx.Header().Set("Content-Type", "application/json")
		ctx.Write(content)
	}
}
