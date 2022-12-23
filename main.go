package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Teste")
	http.HandleFunc("/", serverTesting)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func serverTesting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", r.URL.Path[1:])
}
