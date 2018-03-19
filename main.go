package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/snamber/serve/middleware"
)

var (
	dir           string
	path          string
	port          string
	basicAuthFlag bool
)

func main() {
	flag.StringVar(&port, "port", "8080", "port on which to serve")
	flag.StringVar(&path, "path", "/", "path on which files will be served")
	flag.StringVar(&dir, "dir", ".", "directory to be served")
	flag.BoolVar(&basicAuthFlag, "basicauth", false, "turn on basic auth")
	flag.Parse()

	r := mux.NewRouter()

	fs := http.FileServer(http.Dir(dir))
	if path != "/" {
		fs = http.StripPrefix(path, fs)
	}

	switch basicAuthFlag {
	case true:
		r.PathPrefix(path).Handler(
			mw.Chain(fs.ServeHTTP,
				mw.BasicAuth(),
				mw.Logging(),
			),
		)
	default:
		r.PathPrefix(path).Handler(
			mw.Chain(fs.ServeHTTP,
				mw.Logging(),
			),
		)
	}

	log.Println("serving", http.Dir(dir), "on localhost:"+port+path)

	http.ListenAndServe(":"+port, r)
}
