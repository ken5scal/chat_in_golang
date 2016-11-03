package main

import "net/http"

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		// Not Authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	} else if err != nil {
		// Error
		panic(err.Error())
		return
	}
	// Call Wrapped handler
	h.next.ServeHTTP(w, r)
}

// MustAuth adapts handler to ensure authentication has occurred.
func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}
