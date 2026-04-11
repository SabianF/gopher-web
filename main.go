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
		data_repos.GetRootPath() + "/features/core/presentation/templates/components/layout.html",
		data_repos.GetRootPath() + "/features/core/presentation/templates/pages/home.html",
	))

	pages["not found"] = template.Must(template.ParseFiles(
		data_repos.GetRootPath() + "/features/core/presentation/templates/components/layout.html",
		data_repos.GetRootPath() + "/features/core/presentation/templates/pages/not_found.html",
	))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	err := pages["home"].ExecuteTemplate(w, "layout", nil)
	if (err != nil) {
		fmt.Printf("%v", err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Url string
	} {
		Url: r.URL.String(),
	}

	err := pages["not found"].ExecuteTemplate(w, "layout", data)
	if (err != nil) {
		fmt.Printf("%v", err)
	}
}

func main() {
	core_pub_dir := http.Dir("./features/core/data/sources/public")
	core_fs := http.FileServer(core_pub_dir)
	core_fs_stripped := http.StripPrefix("/features/core/data/sources/public/", core_fs)
	http.Handle("/features/core/data/sources/public/", core_fs_stripped)

	http.HandleFunc("/dashboard", rootHandler)
	http.HandleFunc("/", notFoundHandler)

	PORT := 3000
	fmt.Printf("Running server on port %d\n", PORT)
	log.Fatal(
		http.ListenAndServe(":" + strconv.Itoa(PORT), nil),
	)
}
