package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/fsnotify/fsnotify"
)

const httpPort = "3000"
const httpsPort = "3001"
const graphqlPort = "8081"

var custom404 string
var highPriorityManifest []Source
var serveMux *http.ServeMux

func main() {
	var err error
	custom404, err = load404()
	if err != nil {
		panic(err.Error())
	}

	watcher := setupFilewatch()
	defer watcher.Close()

	serveMux = http.NewServeMux()
	secureServer := &http.Server{
		Addr:      ":" + httpsPort,
		Handler:   serveMux,
		TLSConfig: &tls.Config{},
	}

	insecureServer := &http.Server{
		Addr:    ":" + httpPort,
		Handler: http.HandlerFunc(redirectToHTTPS),
	}
	setupRoutes()

	go secureServer.ListenAndServeTLS("cert.pem", "key.pem")
	insecureServer.ListenAndServe()
	done := make(chan struct{})
	<-done
	log.Printf("Time to go")
}

func handle404(resp http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(resp, custom404)
}

func push(resp http.ResponseWriter) {
	p, ok := resp.(http.Pusher)
	if !ok {
		return
	}

	for _, v := range highPriorityManifest {
		p.Push(v.Source, nil)
	}

	for _, v := range highPriorityManifest {
		if v.CSS != nil {
			p.Push(*v.CSS, nil)
		}
	}
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

func setupFilewatch() *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
					setupRoutes()
				}
			case watchErr := <-watcher.Errors:
				log.Println("watch error:", watchErr)
			}
		}
	}()

	err = watcher.Add("./resources/manifest.json")
	if err != nil {
		log.Fatal(err)
	}

	return watcher
}

func setupRoutes() {
	_, highPriorityManifest = loadManifest()

	index, err := loadTemplates(highPriorityManifest)
	if err != nil {
		panic(err.Error())
	}

	serveMux.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

	serveMux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		if req.URL.RequestURI() != "/" {
			// These are not the droids you're looking for.
			handle404(resp, req)
			return
		}
		push(resp)
		fmt.Fprint(resp, index)
	})

	// http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	// http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
	// 	if req.URL.RequestURI() != "/" {
	// 		// These are not the droids you're looking for.
	// 		handle404(resp, req)
	// 		return
	// 	}
	// 	push(resp)
	// 	fmt.Fprint(resp, index)
	// })
}
