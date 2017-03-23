package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

const httpPort = "3000"
const httpsPort = "3001"
const graphqlPort = "8081"

func main() {
	rawJSON, err := ioutil.ReadFile("./resources/manifest.json")
	if err != nil {
		panic(err.Error())
	}

	var manifest map[string]Source
	err = json.Unmarshal(rawJSON, &manifest)
	if err != nil {
		panic(err.Error())
	}

	manifestSlice := make([]Source, len(manifest))
	offset, priorityCount := 0, 0
	for _, v := range manifest {
		manifestSlice[offset] = v
		if v.Priority != nil {
			priorityCount++
		}
		offset++
	}
	sort.Slice(manifestSlice, func(i int, j int) bool {
		if manifestSlice[i].Priority == nil {
			return false
		}
		if manifestSlice[j].Priority == nil {
			return true
		}
		return *manifestSlice[i].Priority < *manifestSlice[j].Priority
	})
	highPriorityManifest := manifestSlice[0:priorityCount]

	template, err := loadTemplates()
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
	fmt.Println(t)

	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	http.HandleFunc("/", func(resp http.ResponseWriter, _ *http.Request) {
		p, ok := resp.(http.Pusher)
		if ok {
			p.Push("/resources/"+manifest["manifest"].Source, nil)
			p.Push("/resources/"+manifest["vendor"].Source, nil)
			p.Push("/resources/"+manifest["app"].Source, nil)
			if manifest["app"].CSS != nil {
				p.Push("/resources/"+*manifest["app"].CSS, nil)
			}
		}
		fmt.Fprint(resp, t)
	})
	go http.ListenAndServeTLS(":"+httpsPort, "cert.pem", "key.pem", nil)
	http.ListenAndServe(":"+httpPort, http.HandlerFunc(redirectToHTTPS))
}

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:8081" will only
	// work if you are accessing the server from your local machine.
	host := r.Host
	offset := strings.IndexByte(host, ':')
	if offset != -1 {
		host = host[0:offset]
	}
	dest := fmt.Sprintf("https://%s:%s%s", host, httpsPort, r.RequestURI)
	http.Redirect(w, r, dest, http.StatusMovedPermanently)
}
