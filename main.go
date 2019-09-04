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
	urlPrefix = "/"
	port      = "8080"
	basicAuth string
	fallback  = ""
)

func main() {

	flaggy.String(&port, "", "port", fmt.Sprintf("port on which to serve (default '%v')", port))
	flaggy.String(&urlPrefix, "", "path", fmt.Sprintf("URLprefix on which files will be served (default '%v')", urlPrefix))
	flaggy.String(&dir, "", "dir", fmt.Sprintf("directory to be served (default '%v')", dir))
	flaggy.String(&basicAuth, "", "user:pass", "username:password combination for basic auth (default off)")
	flaggy.String(&fallback, "", "fallback", "the file to serve instead of 404s, if desired (default off)")
	flaggy.Parse()

	fs := NewFileServer(dir, urlPrefix, basicAuth, fallback)

	urlPrefix = addSlashesToPath(urlPrefix)
	handler := http.StripPrefix(urlPrefix, fs)

	fmt.Println("listening on", ":"+port+urlPrefix)
	http.ListenAndServe(":"+port, handler)
}

// FileServer is a http.Handler
type FileServer struct {
	handleFunc func(w http.ResponseWriter, r *http.Request)
}

// NewFileServer creates a n FileServer object with a configured handleFunc
func NewFileServer(dir, urlPrefix, basicAuth, fallback string) *FileServer {
	fs := new(FileServer)

	// assemble middleware for Handlefunc according to specification
	handleFunc := http.FileServer(http.Dir(dir)).ServeHTTP
	if fallback != "" {
		handleFunc = middleware.Chain(handleFunc, middleware.Fallback(dir, urlPrefix, fallback))
	}
	if basicAuth != "" {
		basic := strings.Split(basicAuth, ":")
		handleFunc = middleware.Chain(handleFunc, middleware.BasicAuth(basic[0], basic[1]))
	}
	handleFunc = middleware.Chain(handleFunc, middleware.Logging())

	fs.handleFunc = handleFunc

	return fs
}

func (f FileServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.handleFunc(w, r)
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
