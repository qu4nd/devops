package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	title := "Jenkins X golang http example"

	if value := os.Getenv("TITLE"); value != "" {
		title = value
	}
	
	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	fmt.Fprintf(w, "Hello from:  "+title+"\n")
	fmt.Fprintf(w, "I like big butts and I cannot lie...\n")
	
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
