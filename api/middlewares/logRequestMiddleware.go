package middlewares

import (
	"log"
	"net/http"
)

type LogRequestMiddleware struct {
}

func NewLogRequestMiddleware() *LogRequestMiddleware {
	return &LogRequestMiddleware{}
}

func (v *LogRequestMiddleware) Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}