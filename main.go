package main

import (
	"github.com/russross/blackfriday"
	"os"

	"net/http"
)

func main() {
	port := os.GetEnv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":"+port, nil)
}

func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
