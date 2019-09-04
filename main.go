package main

import (
	"fmt"
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

	flaggy.String(&port, "", "port", "port on which to serve (default 8080)")
	flaggy.String(&path, "", "path", "path on which files will be served (default \"/\")")
	flaggy.String(&dir, "", "dir", "directory to be served (default \".\")")
	flaggy.String(&basicAuth, "", "user:pass", "username:password combination for basic auth (default off)")
	flaggy.Parse()

	path = addSlashesToPath(path)
	fs := http.StripPrefix(path, http.FileServer(http.Dir(dir)))

	handleFunc := middleware.Chain(fs.ServeHTTP,
		middleware.Logging(),
	)
	if basicAuth != "" {
		basic := strings.Split(basicAuth, ":")
		handleFunc = middleware.Chain(fs.ServeHTTP,
			middleware.BasicAuth(basic[0], basic[1]),
			middleware.Logging())
	}

	http.HandleFunc(path, handleFunc)
	fmt.Println("listening on", ":"+port+path)
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
