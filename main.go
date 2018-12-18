package main

import (
	"flag"
	"net/http"
	"strings"
)

var (
	dir  string
	path string
	port string
)

func main() {
	flag.StringVar(&port, "port", "8080", "port on which to serve")
	flag.StringVar(&path, "path", "/", "path on which files will be served")
	flag.StringVar(&dir, "dir", ".", "directory to be served")
	flag.Parse()

	path = addSlashesToPath(path)

	http.Handle(path, http.StripPrefix(path, http.FileServer(http.Dir(dir))))
	http.ListenAndServe(":"+port, nil)
}

func addSlashesToPath(path string) string {
	path = strings.Trim(path, "/")
	if path == "" {
		path = "/"
	} else {
		path = "/" + path + "/"
	}
	return path
}
