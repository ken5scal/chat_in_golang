package main

import (
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/labstack/gommon/log"
)

type templateHandler struct {
	once    sync.Once
	filenae string
	templ   *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Lazy Initialization.
	t.once.Do(func() {
		// once.Do only executes function once (even multiple goroutine calls serveHTTP\
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filenae)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	r := newRoom()
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte(`<html><body>hogehoge</body></html>`))
	//})
	http.Handle("/", &templateHandler{filenae: "chat.html"})
	http.Handle("/", r)
	go r.run()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
