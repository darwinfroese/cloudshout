package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type index struct {
	Title string
}

type admin struct {
	Title     string
	Heading   string
	Templates []string
}

type theme struct {
	Name            string
	FontColor       string
	BackgroundColor string
}

// RenderIndex - Will be used to render our index.html web page
func RenderIndex(w http.ResponseWriter, r *http.Request) {
	i := index{Title: "Cloudshout Official Blog"}

	if pusher, ok := w.(http.Pusher); ok {
		if err := pusher.Push("/main.css", nil); err != nil {
			fmt.Println("Error:", err.Error())
		}
	} else {
		fmt.Println("Couldn't push main.css")
	}

	t, err := template.ParseFiles("web/index.html")

	if err != nil {
		fmt.Println("Error:", err.Error())

		e := http.StatusInternalServerError
		http.Error(w, http.StatusText(e), e)
		return
	}

	t.Execute(w, i)
}

// RenderCSS - Serves main.css
func RenderCSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/main.css")
}

// RenderAdmin - Will render the admin page with templated options
func RenderAdmin(w http.ResponseWriter, r *http.Request) {
	a := admin{Title: "Cloudshout Official Blog - Admin", Heading: "Admin Page"}
	a.Templates = []string{"Text Post", "Half Width Text Post"}

	if pusher, ok := w.(http.Pusher); ok {
		if err := pusher.Push("/main.css", nil); err != nil {
			fmt.Println("Error:", err.Error())
		}
	} else {
		fmt.Println("Couldn't push main.css")
	}

	t, err := template.ParseFiles("web/admin.html")

	if err != nil {
		fmt.Println("Error:", err.Error())

		e := http.StatusInternalServerError
		http.Error(w, http.StatusText(e), e)
		return
	}

	t.Execute(w, a)
}

// CreatePostHandler - Creates a blog post
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Create Blog Post
	}
}
