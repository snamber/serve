package main

import (
	"flag"
	"log"
	"net/http"
)

var dir string

func main() {

	flag.StringVar(&dir, "dir", ".", "directory to be served")
	flag.Parse()

	fs := http.FileServer(http.Dir(dir))

	log.Println("serving", http.Dir(dir), "on localhost:8080")
	http.ListenAndServe(":8080", fs)
}
