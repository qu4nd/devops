package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	title := "Jenkins X golang http example"

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	fmt.Fprintf(w, "Hello from:  "+title+"\n")
	fmt.Fprintf(w, "I like big butts and I cannot lie...\n")
	
	ticker := time.NewTicker(5000 * time.Millisecond)
    	go func() {
        for t := range ticker.C {
            fmt.Println("Tick at", t)
        	}
    	}()
	time.Sleep(1600 * time.Millisecond)
    	ticker.Stop()
    	fmt.Println("Ticker stopped")
	
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
