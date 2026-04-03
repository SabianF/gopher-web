package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	html_templating "github.com/SabianF/gopher-web/features/core/data/repos"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	root_page := html_templating.ParseFiles(nil, "./features/core/presentation/templates/layout.html")
	w.Write(root_page.Bytes())
}

func main() {
	http.HandleFunc("/", rootHandler)

	PORT := 3000
	fmt.Printf("Running server on port %d\n", PORT)
	log.Fatal(
		http.ListenAndServe(":" + strconv.Itoa(PORT), nil),
	)
}
