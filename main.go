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
	root_dir := data_repos.GetRootPath()
	core_pages_dir := root_dir + "/features/core/presentation/templates/pages/"
	layout_path := root_dir + "/features/core/presentation/templates/components/layout.html"

	pages = make(map[string]*template.Template)

	pages["home"] = template.Must(template.ParseFiles(
		layout_path,
		core_pages_dir + "home.html",
	))

	pages["dashboard"] = template.Must(template.ParseFiles(
		layout_path,
		core_pages_dir + "dashboard.html",
	))

	pages["not found"] = template.Must(template.ParseFiles(
		layout_path,
		core_pages_dir + "not_found.html",
	))
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	err := pages["dashboard"].ExecuteTemplate(w, "layout", nil)
	if (err != nil) {
		fmt.Printf("%v", err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Url string
	} {
		Url: r.URL.String(),
	}

	if (data.Url != "/") {
		w.WriteHeader(http.StatusNotFound)
		err := pages["not found"].ExecuteTemplate(w, "layout", data)
		if (err != nil) {
			fmt.Printf("%v", err)
		}
		return
	}

	err := pages["home"].ExecuteTemplate(w, "layout", data)
	if (err != nil) {
		fmt.Printf("%v", err)
	}
}

func main() {
	core_pub_dir := http.Dir("./features/core/data/sources/public")
	core_fs := http.FileServer(core_pub_dir)
	core_fs_stripped := http.StripPrefix("/public/core/", core_fs)
	http.Handle("/public/core/", core_fs_stripped)

	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/dashboard", handleDashboard)

	PORT := 3000
	fmt.Printf("Running server on port %d\n", PORT)
	log.Fatal(
		http.ListenAndServe(":" + strconv.Itoa(PORT), nil),
	)
}
