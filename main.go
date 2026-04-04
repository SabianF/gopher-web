package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/SabianF/gopher-web/features/core/presentation/templates"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	layout_data := templates.LayoutData{
		Title: "The Title",
		MetaDescription: "This is the meta description.",
		PageBody: "<p>Page body.</p>",
	}

	root_page := templates.Layout(layout_data)

	w.Write([]byte(root_page))
}

func main() {
	http.HandleFunc("/", rootHandler)

	PORT := 3000
	fmt.Printf("Running server on port %d\n", PORT)
	log.Fatal(
		http.ListenAndServe(":" + strconv.Itoa(PORT), nil),
	)
}
