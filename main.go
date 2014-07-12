package main

import (
	"sync"
	"log"
	"net/http"
)

type doc struct {
	_ sync.Mutex
	dirty bool
	text []byte
}

func main(){
	editor := docEditor{}
	http.HandleFunc("/edit/", editor.editHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type docEditor struct{
	docs map[string]doc

}
func (ed *docEditor) editHandle(rw http.ResponseWriter ,req *http.Request){
	docId := req.URL.Path[len("/edit/"):]
	text := ed.getDoc(docId)
	_ = text
	rw.Write()
	
}
func(ed *docEditor) getDoc(id string) []byte{
	return ed.docs[id].text
	
}
