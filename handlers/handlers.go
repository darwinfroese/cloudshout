package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type index struct {
	Title string
	Posts []post
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

type post struct {
	key         int
	Title       string
	Description string
	Template    string
	Post        string
	Body        template.HTML
	URL         string
}

var posts []post
var uniqueID = 1000
var templateKeys = []string{"Text Post", "Half Width Text Post"}
var templateMap = map[string]string{
	"Text Post":            "web/templates/full_post.html",
	"Half Width Text Post": "web/templates/half_post.html",
}

// RenderIndex - Will be used to render our index.html web page
func RenderIndex(w http.ResponseWriter, r *http.Request) {
	i := index{Title: "Cloudshout Official Blog", Posts: posts}

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

// ServeJS - Serves main.js
func ServeJS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/main.js")
}

// RenderAdmin - Will render the admin page with templated options
func RenderAdmin(w http.ResponseWriter, r *http.Request) {
	a := admin{Title: "Cloudshout Official Blog - Admin", Heading: "Admin Page"}
	a.Templates = templateKeys

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

// RenderPost - renders the post based on the key passed in as a var
func RenderPost(w http.ResponseWriter, r *http.Request) {
	key, _ := strconv.Atoi(r.URL.Query()["key"][0])

	var p post

	for _, pp := range posts {
		if pp.key == key {
			p = pp
		}
	}

	// check if p is empty
	if (post{}) == p {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// serve template with p
	f := templateMap[p.Template]
	t, err := template.ParseFiles("web/templates/post.html", f)

	if err != nil {
		fmt.Println("Error:", err.Error())

		e := http.StatusInternalServerError
		http.Error(w, http.StatusText(e), e)
		return
	}

	t.Execute(w, p)
}

// CreatePostHandler - Creates a blog post
func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var p post
		json.NewDecoder(r.Body).Decode(&p)

		p.key = uniqueID
		p.URL = "/blog?key=" + strconv.Itoa(uniqueID)
		uniqueID++

		p.Post = strings.Replace(p.Post, "\n", "<br>", -1)
		p.Body = template.HTML(p.Post)

		posts = append(posts, p)
	}
}
