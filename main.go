package main

import (
	"log"
	"net/http"
)

func main() {
	editor := initEditor()
	editor.docs["10101001"] = doc{Text: []byte("The  quick brown fox jumped \x0A over the lazy dog \xCF\x80 \n 	=================")}
	http.HandleFunc("/view/", editor.viewHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
