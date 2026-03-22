package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request for " + r.URL.Path)
	w.WriteHeader(200)
}

func main() {
	http.HandleFunc("/", rootHandler)

	PORT := 3000
	fmt.Printf("Running server on port %d\n", PORT)
	log.Fatal(
		http.ListenAndServe(":" + strconv.Itoa(PORT), nil),
	)
}
