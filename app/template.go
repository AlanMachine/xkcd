package app

import (
	"os"
	"strings"
	"text/template"
)

const txtTempl = `{{range .}}
URL: https://xkcd.com/{{.Num}}/
Title: {{.Title}}
{{end}}`

const htmlTempl = `{{range .}}
<h1>{{.Title}}</h1>
<p>{{.Day | trim | printf "%02s"}}.{{.Month | trim | printf "%02s"}}.{{.Year}}</p>
<p><img src="{{.Img}}"/></p>
<p>{{.Transcript}}</p>
<hr/>
{{end}}`

func formatOutput(wc []webComic, mode int) error {
	templType := txtTempl
	if mode == 1 {
		templType = htmlTempl
	}

	templ := template.Must(template.New("webcomic").
		Funcs(template.FuncMap{"trim": trimSpaces}).
		Parse(templType))
	if err := templ.Execute(os.Stdout, wc); err != nil {
		return err
	}

	return nil
}

func trimSpaces(s string) string {
	return strings.Trim(s, " \t\r\n")
}
