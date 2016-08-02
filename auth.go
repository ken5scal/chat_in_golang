package main

import (
	"net/http"
	"strings"
	"log"
	"fmt"
	"github.com/stretchr/gomniauth"
)

type authHandler struct {
	next http.Handler // handler which will be handled
}

func (handler *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := r.Cookie("auth"); err == http.ErrNoCookie {
		// not authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error())
	} else {
		// succeeded
		handler.next.ServeHTTP(w, r)
	}
}

func MustAuth(handler http.Handler) http.Handler{
	return &authHandler{next: handler}
}

// loginHandlerは3rd partyへのログイン処理を受け持つ
// loginHandler managers 3rd party authentication process
// Path Format: /auth/{action}/{provider}
func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":
		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("Failed fetching auth provider: ", provider, "-", err)
		}
		loginUrl, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			log.Fatalln("Failed when calling GetBeginAuthURL: ", provider, "-", err)
		}
		w.Header().Set("Location", loginUrl)
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound) //404
		fmt.Fprintf(w, "cannot process action %s", action)
	}
}