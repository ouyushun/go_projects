package middleware

import (
	"fmt"
	"net/http"
)

type LoginverificationMiddleware struct {
}

func NewLoginverificationMiddleware() *LoginverificationMiddleware {
	return &LoginverificationMiddleware{}
}

func (m *LoginverificationMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		fmt.Println("LoginVerificationMiddleware")

		// Passthrough to next handler if need
		next(w, r)
	}
}
