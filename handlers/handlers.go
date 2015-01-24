package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gophergala/serradacapivara/db"
	"github.com/zenazn/goji/web"
)

// Index is the website homepage
func Index(c web.C, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html.go")

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

// Search is the page where the results of search are shown
func Search(c web.C, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/search.html.go")

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}

// Map is the page where all sites are shown at once
func Map(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Map")
}

// Site is the page that describe the given site
func Site(c web.C, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/site.html.go")

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	site, err := db.FindByID(c.URLParams["id"])

	if err != nil {
		log.Println(err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return

	}

	t.Execute(w, site)
}

// About the project
func About(c web.C, w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/about.html.go")

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	t.Execute(w, nil)
}
