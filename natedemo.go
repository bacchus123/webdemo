package main

import (
	"fmt"
	_ "github.com/gorilla/websocket"
	"html/template"
	"net/http"
	"sync"
)

type doc struct {
	_     sync.Mutex
	dirty bool
	Text  []byte
}

type docEditor struct {
	docs map[string]doc
	view *template.Template
}

func initEditor() *docEditor {
	ed := docEditor{}
	ed.docs = make(map[string]doc)
	ed.view = template.Must(template.ParseFiles("templates/view.html"))
	return &ed
}

func (ed *docEditor) getDocText(id string) []byte {
	return ed.docs[id].Text

}

func (ed *docEditor) editHandle(rw http.ResponseWriter, req *http.Request) {
	docId := req.URL.Path[len("/edit/"):]
	text := ed.getDocText(docId)
	_ = text
	fmt.Fprint(rw, text)

}

func (ed *docEditor) viewHandle(rw http.ResponseWriter, req *http.Request) {
	docId := req.URL.Path[len("/view/"):]
	requestedDoc := ed.docs[docId]
	err := ed.view.Execute(rw, requestedDoc)
	if err != nil {
		panic(err)
	}
}
