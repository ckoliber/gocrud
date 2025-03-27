package repository

import (
	"text/template"
)

var postgres *template.Template

const postgresTpl = `
{{ define "create" }}
{{ end }}
 aaaa
`

func getPostgres() *template.Template {
	if postgres == nil {
		tpl, err := template.New("postgres").Parse(postgresTpl)
		if err != nil {
			panic(err)
		} else {
			postgres = tpl
		}
	}

	return postgres
}
