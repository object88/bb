package main

import (
	"html/template"
	"io/ioutil"
)

func loadTemplates() (*template.Template, error) {
	indexTemplate, _ := ioutil.ReadFile("./templates/index.tmpl")
	scriptTemplate, _ := ioutil.ReadFile("./templates/script.tmpl")
	t := template.New("index")
	t, _ = t.Parse(string(indexTemplate))
	t, _ = t.Parse(string(scriptTemplate))

	return t, nil
}
