package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request for " + r.URL.Path)
	root_page, err := os.ReadFile("features/core/presentation/templates/layout.html")
	if (err != nil) {
		panic(err)
	}
	w.Write(root_page)
}

func main() {
	http.HandleFunc("/", rootHandler)

	PORT := 3000
	fmt.Printf("Running server on port %d\n", PORT)
	log.Fatal(
		http.ListenAndServe(":" + strconv.Itoa(PORT), nil),
	)
}
