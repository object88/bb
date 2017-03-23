package main

import (
	"html/template"
	"io/ioutil"
)

func loadTemplates() (*template.Template, error) {
	indexTemplate, _ := ioutil.ReadFile("./templates/index.tmpl")
	scriptTemplate, _ := ioutil.ReadFile("./templates/script.tmpl")
	stylesheetTemplate, _ := ioutil.ReadFile("./templates/stylesheet.tmpl")
	t := template.New("index")
	t, _ = t.Parse(string(indexTemplate))
	t, _ = t.Parse(string(scriptTemplate))
	t, _ = t.Parse(string(stylesheetTemplate))

	return t, nil
}
