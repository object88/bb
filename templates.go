package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

func loadTemplates(highPriorityManifest []Source) (string, error) {
	template, err := readTemplates()
	if err != nil {
		panic(err.Error())
	}

	var buf bytes.Buffer
	foo := struct {
		APIKey   string
		ClientID string
		Scripts  []Source
	}{
		APIKey:   "123",
		ClientID: "456",
		Scripts:  highPriorityManifest,
	}
	err = template.Execute(&buf, foo)
	if err != nil {
		panic(err.Error())
	}
	t := buf.String()

	return t, nil
}

func readTemplates() (*template.Template, error) {
	indexTemplate, _ := ioutil.ReadFile("./templates/index.tmpl")
	scriptTemplate, _ := ioutil.ReadFile("./templates/script.tmpl")
	stylesheetTemplate, _ := ioutil.ReadFile("./templates/stylesheet.tmpl")
	t := template.New("index")
	t, _ = t.Parse(string(indexTemplate))
	t, _ = t.Parse(string(scriptTemplate))
	t, _ = t.Parse(string(stylesheetTemplate))

	return t, nil
}
