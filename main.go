package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var dir string
var path string
var port string

func main() {
	flag.StringVar(&port, "port", "8080", "port on which to serve")
	flag.StringVar(&path, "path", "/", "path on which files will be served")
	flag.StringVar(&dir, "dir", ".", "directory to be served")
	flag.Parse()

	r := mux.NewRouter()

	// fileserver
	fs := http.FileServer(http.Dir(dir))
	if path != "/" {
		fs = http.StripPrefix(path, fs)
	}

	r.PathPrefix(path).Handler(fs)

	log.Println("serving", http.Dir(dir), "on localhost:"+port+path)
	http.ListenAndServe(":"+port, r)
}
