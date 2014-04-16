package her

import (
	"html/template"
	"runtime"
)

var tpl = `
<!doctype html>
<html>
<head>
	<meta charset="UTF-8">
	<title>Application Error</title>
	<style></style>
</head>
<body>
	<div id="header">
		<h2></h2>
	</div>
	<div id="content"></div>
	<div id="footer">
		her version {{HerVersion}} , go version {{GoVersion}}
	</div>
</body>
</html>
`

func WriteError(ctx *Context) {
	t, _ := template.New("her-error.html").Parse(tpl)
	data := make(map[string]string)
	data["HerVersion"] = Version
	data["GoVersion"] = runtime.Version()
	ctx.WriteHeader(500)
	t.Execute(ctx.ResponseWriter, data)
}
