package main

import (
	"net/http"
	"text/template"
	"sync"
	"github.com/labstack/gommon/log"
	"path/filepath"
)

type templateHandler struct {
	once    sync.Once
	filenae string
	templ   *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filenae)))
	})
	t.templ.Execute(w, nil)
}

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte(`<html><body>hogehoge</body></html>`))
	//})
	http.Handle("/", &templateHandler{filenae: "chat.html"})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
