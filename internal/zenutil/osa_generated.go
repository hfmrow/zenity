// Code generated by zenity; DO NOT EDIT.
// +build darwin

package zenutil

import "encoding/json"
import "text/template"

var scripts = template.Must(template.New("").Funcs(template.FuncMap{"json": func(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}}).Parse(`
{{define "color" -}}tell application (path to frontmost application as text)
activate
{{if .Color -}}
set c to choose color default color { {{index .Color 0}},{{index .Color 1}},{{index .Color 2}} }
{{else -}}
set c to choose color
{{end}}
"rgb(" & (item 1 of c) div 256 & "," & (item 2 of c) div 256 & "," & (item 3 of c) div 256 & ")"
end tell
{{- end}}
{{define "file" -}}var app = Application.currentApplication()
app.includeStandardAdditions = true
app.activate()
var opts = {}
{{if .Prompt -}}
opts.withPrompt = {{json .Prompt}}
{{end -}}
{{if .Type -}}
opts.ofType = {{json .Type}}
{{end -}}
{{if .Name -}}
opts.defaultName = {{json .Name}}
{{end -}}
{{if .Location -}}
opts.defaultLocation = {{json .Location}}
{{end -}}
{{if .Invisibles -}}
opts.invisibles = {{json .Invisibles}}
{{end -}}
{{if .Multiple -}}
opts.multipleSelectionsAllowed = {{json .Multiple}}
{{end -}}
var res = app[{{json .Operation}}](opts)
if (Array.isArray(res)) {
res.join({{json .Separator}})
} else {
res.toString()
}
{{- end}}
{{define "msg" -}}var app = Application.currentApplication()
app.includeStandardAdditions = true
app.activate()
var opts = {}
{{if .Message -}}
opts.message = {{json .Message}}
{{end -}}
{{if .As -}}
opts.as = {{json .As}}
{{end -}}
{{if .Title -}}
opts.withTitle = {{json .Title}}
{{end -}}
{{if .Icon -}}
opts.withIcon = {{json .Icon}}
{{end -}}
{{if .Buttons -}}
opts.buttons = {{json .Buttons}}
{{end -}}
{{if .Cancel -}}
opts.cancelButton = {{json .Cancel}}
{{end -}}
{{if .Default -}}
opts.defaultButton = {{json .Default}}
{{end -}}
{{if .Timeout -}}
opts.givingUpAfter = {{json .Timeout}}
{{end -}}
var res = app[{{json .Operation}}]({{json .Text}}, opts)
if (res.gaveUp) {
ObjC.import("stdlib")
$.exit(5)
}
if (res.buttonReturned === {{json .Extra}}) {
res
} else {
void 0
}
{{- end}}
{{define "notify" -}}var app = Application.currentApplication()
app.includeStandardAdditions = true
app.activate()
var opts = {}
{{if .Title -}}
opts.withTitle = {{json .Title}}
{{end -}}
{{if .Subtitle -}}
opts.subtitle = {{json .Subtitle}}
{{end -}}
void app.displayNotification({{json .Text}}, opts)
{{- end}}`))
