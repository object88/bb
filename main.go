package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

const httpPort = 3000
const httpsPort = 3001
const graphqlPort = 8081

func main() {
	template, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err.Error())
	}

	var buf bytes.Buffer
	foo := struct {
		APIKey   string
		ClientID string
	}{
		APIKey:   "123",
		ClientID: "456",
	}
	err = template.ExecuteTemplate(&buf, "index.html", foo)
	if err != nil {
		panic(err.Error())
	}
	t := buf.String()

	url, _ := url.Parse(fmt.Sprintf("https://localhost:%d/graphql", graphqlPort))
	fmt.Printf("Setting up proxy for %s\n", url.String())
	proxy := httputil.NewSingleHostReverseProxy(url)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/graphql/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Printf("Proxying %s\n", req.URL.String())
		proxy.ServeHTTP(resp, req)
	})
	http.HandleFunc("/", func(resp http.ResponseWriter, _ *http.Request) {
		p, ok := resp.(http.Pusher)
		if ok {
			p.Push("/public/app.bundle.js", nil)
			p.Push("/public/css/app.css", nil)
		}
		fmt.Fprint(resp, t)
	})
	go http.ListenAndServeTLS(fmt.Sprintf(":%d", httpsPort), "cert.pem", "key.pem", nil)
	http.ListenAndServe(fmt.Sprintf(":%d", httpPort), http.HandlerFunc(redirectToHTTPS))
}

func redirectToHTTPS(w http.ResponseWriter, r *http.Request) {
	// Redirect the incoming HTTP request. Note that "127.0.0.1:8081" will only
	// work if you are accessing the server from your local machine.
	host := r.Host
	offset := strings.IndexByte(host, ':')
	if offset != -1 {
		host = host[0:offset]
	}
	dest := fmt.Sprintf("https://%s:%d%s", host, httpsPort, r.RequestURI)
	http.Redirect(w, r, dest, http.StatusMovedPermanently)
}
