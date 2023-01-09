package middleware

import (
	"context"
	"log"
	"net/http"
)

func OnlyAuthorised(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Haha " + r.Method)
		next.ServeHTTP(w, r)
	})
}

func BlaBla(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Blablabla")
		ctx := r.Context()
		ctx = context.WithValue(ctx, "bla", "bla")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
