package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {

	target, err := url.Parse("http://localhost:9090")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	http.Handle("/echo", proxy)
	http.Handle("/banaan", proxy)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte(`HEY`))
	})

	log.Fatalln(http.ListenAndServe("localhost:8080", nil))
}
