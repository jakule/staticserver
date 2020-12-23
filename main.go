package main

import (
	"embed"
	"log"
	"net/http"
)

//go:embed *
var files embed.FS

func main() {
	http.Handle("/", http.FileServer(http.FS(files)))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
