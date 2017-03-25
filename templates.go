package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

func loadTemplates(highPriorityManifest []Source) (string, error) {
	template, err := readTemplates("index", "./templates/index.tmpl", "./templates/script.tmpl", "./templates/stylesheet.tmpl")
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

func load404() (string, error) {
	template, err := readTemplates("404", "./templates/404.tmpl")
	if err != nil {
		panic(err.Error())
	}

	var buf bytes.Buffer
	err = template.Execute(&buf, nil)
	t := buf.String()
	return t, nil
}

func readTemplates(templateName string, fileNames ...string) (*template.Template, error) {
	t := template.New(templateName)
	for _, v := range fileNames {
		file, _ := ioutil.ReadFile(v)
		t, _ = t.Parse(string(file))
	}

	return t, nil
}
