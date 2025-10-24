package main

import (
"fmt"
"log"
"net/http"
"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    // Avoid logging raw query parameters to prevent leaking sensitive data
    log.Printf("method=%s path=%q remote=%s ua=%q", r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
    fmt.Fprintf(w, `
          ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/

	
Hello from Docker!

`)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

