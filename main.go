package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template
func init() {
	var err error
	tpl, err = tpl.ParseGlob("templates/*.gohtml")

	if err != nil {
		log.Panicln(err)
	}
}

func main() {
	fs :=  http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	http.HandleFunc("/", homePageRoute)
	http.HandleFunc("/post/", postPageRoute)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

type post struct{Title string; URL string}
type postListing []post
type pageData struct {
	Title string;
	Posts postListing;
}

func homePageRoute(w http.ResponseWriter, r *http.Request){
	data := pageData{
		Title: "Code of Fury - Index",
		Posts: postListing{
			post{Title: "Hello post", URL: "/post/hello"},
			post{Title: "About go", URL: "/post/about-go"},
		},
	}
	err := tpl.ExecuteTemplate(w, "index.gohtml", data)

	if err != nil {
		log.Panicln(err)
	}
}

func postPageRoute(w http.ResponseWriter, r *http.Request){
	postID := r.URL.Path[len("/post/"):]
	fmt.Println(postID)
	data := struct{
		Title string;
		Body string;
	}{
		Title: "Go related blog",
		Body: "Hellow from text here",
	}
	err := tpl.ExecuteTemplate(w, "post.gohtml", data)

	if err != nil {
		log.Panicln(err)
	}
}
