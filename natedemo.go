package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"net/http"
	"sync"
	"time"
)

type doc struct {
	_    sync.Mutex
	sent []bool
	Text []byte
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

	conn, err := upgrader.Upgrade(rw, req, nil)
	docId := req.URL.Path[len("/edit/"):]

}

var upgrader = websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}

func (ed *docEditor) viewHandle(rw http.ResponseWriter, req *http.Request) {
	docId := req.URL.Path[len("/view/"):]
	requestedDoc := ed.docs[docId]
	err := ed.view.Execute(rw, requestedDoc)
	if err != nil {
		panic(err)
	}

	conn, err := websocket.upgrader.Upgrade(rw, req, nil)
	if err != nil {
		panic(err)
	}
	go func() {
		_, _, err = conn.ReadMessage()
		if err != null {
			conn.Close()
			return
		}
	}()
	for {
		time.Sleep(17 * time.Millisecond) // 1/60th of a second
		err := conn.WriteMessage(websocket.BinaryMessage, requestedDoc.Text)
		if err != nil {
			conn.Close()
			return
		}

	}

}
