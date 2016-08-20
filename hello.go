package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	h, _ := os.Hostname()

	fmt.Fprintf(w, "Hi there, I'm Version 0.2 served from %s!", h)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8484", nil)
}

// only used by unit test
func runLocal() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	return mux
}
