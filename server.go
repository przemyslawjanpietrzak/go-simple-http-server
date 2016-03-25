package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path = "."
	}
	var port bytes.Buffer

	port.WriteString(":")
	if len(os.Args) > 2 {
		port.WriteString(os.Args[2])
	} else {
		port.WriteString("8080")
	}
	fmt.Println("Serving files in the current directory on port" + port.String() + path)

	http.Handle("/", http.FileServer(http.Dir(path)))
	err := http.ListenAndServe(port.String(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
