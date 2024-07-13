package route

import (
	"net/http"

	h "qualitech.paseto-auth/handler"
	wrapper "qualitech.paseto-auth/middleware"
)


func InitializeRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", h.LoginHandler)
	mux.Handle("/protected", wrapper.Authenticate(http.HandlerFunc(h.ProtectedHandler)))
	mux.HandleFunc("/unprotected", h.UnProtectedHandler)

	return mux
}


