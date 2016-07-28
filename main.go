package main

import (
	"net/http"
	"log"
	"sync"
	"html/template"
	"path/filepath"
)

type templateHandler struct {
	once sync.Once
	filename string
	templ *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	t.templ.Execute(w, nil)
}

func main() {
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte(`
	//	<html>
	//		<head>
	//			<title> Chat </title>
	//		</head>
	//		<body>
	//			Let's Chat
	//		</body>
	//	</html>
	//	`))
	//})
	r := newRoom()
	t1 := templateHandler{filename: "chat.html"}
	http.Handle("/", &t1)
	http.Handle("/room", r)
	go r.run() // start chat room

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal("ListenAndService:", err)
	}
}
