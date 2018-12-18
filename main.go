package main

import (
	"net/http"
	"strings"

	"github.com/integrii/flaggy"

	"github.com/snamber/serve/middleware"
)

var (
	dir       = "."
	path      = "/"
	port      = "8080"
	basicAuth string
)

func main() {

	flaggy.String(&port, "", "port", "port on which to serve")
	flaggy.String(&path, "", "path", "path on which files will be served")
	flaggy.String(&dir, "", "dir", "directory to be served")
	flaggy.String(&basicAuth, "", "user:pass", "username:password combination for basic auth")
	flaggy.Parse()

	path = addSlashesToPath(path)
	fs := http.StripPrefix(path, http.FileServer(http.Dir(dir)))

	handleFunc := mw.Chain(fs.ServeHTTP,
		mw.Logging(),
	)
	if basicAuth != "" {
		basic := strings.Split(basicAuth, ":")
		handleFunc = mw.Chain(fs.ServeHTTP,
			mw.BasicAuth(basic[0], basic[1]),
			mw.Logging())
	}

	http.HandleFunc(path, handleFunc)
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
