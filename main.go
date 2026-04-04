package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	data_repos "github.com/SabianF/gopher-web/features/core/data/repos"
)

var pages map[string]*template.Template

func init() {
	pages = make(map[string]*template.Template)

	pages["home"] = template.Must(template.ParseFiles(
		data_repos.GetRootPath() + "/features/core/presentation/templates/pages/home.html",
		data_repos.GetRootPath() + "/features/core/presentation/templates/components/layout.html",
	))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	err := pages["home"].ExecuteTemplate(w, "layout", nil)
	if (err != nil) {
		fmt.Printf("%v", err)
	}
}

func main() {
	http.HandleFunc("/", rootHandler)

	PORT := 3000
	fmt.Printf("Running server on port %d\n", PORT)
	log.Fatal(
		http.ListenAndServe(":" + strconv.Itoa(PORT), nil),
	)
}
