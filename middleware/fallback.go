package middleware

// Leverages nemo's answer in http://stackoverflow.com/a/21937924/556573
import (
	"net/http"
	"os"
	"path"
)

// Fallback has to be called right before the fileserver
// it serves the fallback-file instead of a 404 in case the
// url-specified path is not actually a file or directory
func Fallback(dir, prefix, file string) Type {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			urlPath := path.Clean(r.URL.Path)
			fullPath := path.Clean(path.Join(dir, urlPath))
			_, err := os.Stat(fullPath)
			if err != nil {
				http.ServeFile(w, r, file)
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}
