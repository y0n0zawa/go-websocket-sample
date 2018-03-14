package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/y0n0zawa/go-websocket-sample/room"
)

type templateHandler struct {
	once     sync.Once
	filename string
	tmpl     *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.tmpl = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})

	t.tmpl.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "Applicationサーバーのアドレス")
	flag.Parse()

	r := room.NewRoom()
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	go r.Run()

	log.Println("Start listening port", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
