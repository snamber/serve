package middleware

// Leverages nemo's answer in http://stackoverflow.com/a/21937924/556573
import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

func BasicAuth(user, pass string) Type {

	// Create a new Middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

			s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
			if len(s) != 2 {
				http.Error(w, "Not authorized", 401)
				return
			}

			b, err := base64.StdEncoding.DecodeString(s[1])
			if err != nil {
				http.Error(w, err.Error(), 401)
				return
			}

			pair := strings.SplitN(string(b), ":", 2)
			if len(pair) != 2 {
				http.Error(w, "Not authorized", 401)
				return
			}

			if pair[0] != user || pair[1] != pass {
				http.Error(w, "Not authorized", 401)
				log.Println("unauthorized access with username", pair[0])
				return
			}

			// Call the next middleware/handler in chain
			f(w, r)
		}
	}
}
