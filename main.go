package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/facebook"
	"github.com/stretchr/gomniauth/providers/github"
	"github.com/stretchr/gomniauth/providers/google"
)

// templ represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP request.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() // parse the flags
	gomniauth.SetSecurityKey("security_key")
	gomniauth.WithProviders(
		facebook.New("1032005158087-2e5glvo84j51kbkg82b3ep77uie1rhr3.apps.googleusercontent.com", "ZC7L37j_Gb08UnwwnixZIzTB", "http://localhost:8080/auth/callback/facebook"),
		github.New("1032005158087-2e5glvo84j51kbkg82b3ep77uie1rhr3.apps.googleusercontent.com", "ZC7L37j_Gb08UnwwnixZIzTB", "http://localhost:8080/auth/callback/github"),
		google.New("1032005158087-2e5glvo84j51kbkg82b3ep77uie1rhr3.apps.googleusercontent.com", "ZC7L37j_Gb08UnwwnixZIzTB", "http://localhost:8080/auth/callback/google"),
	)

	r := newRoom()
	//r.tracer = trace.New(os.Stdout)

	//http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/chat", MustAuth(&templateHandler{filename: "chat.html"}))
	http.Handle("/login", &templateHandler{filename: "login.html"})
	http.HandleFunc("/auth/", loginHandler)
	http.Handle("/room", r)

	// get the room going
	go r.run()

	// start the web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}