package middleware

import (
	"net/http"
)

type Type func(http.HandlerFunc) http.HandlerFunc

// Chain applies middlewares to a http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Type) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}
