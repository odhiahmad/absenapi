package tokenValidation

import (
	"bri-rece/api/middlewares"
	"net/http"
)

type TokenValidationMiddleware struct {
}

func NewTokenValidationMiddleware() *TokenValidationMiddleware  {
	return &TokenValidationMiddleware{}
}

func (v *TokenValidationMiddleware) SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := middlewares.TokenValid(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}