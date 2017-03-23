package main

import (
	"fmt"
	"net/http"
	"strings"
)

const httpPort = "3000"
const httpsPort = "3001"
const graphqlPort = "8081"

func main() {
	manifest, highPriorityManifest := loadManifest()

	index, err := loadTemplates(highPriorityManifest)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(index)

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
		fmt.Fprint(resp, index)
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
