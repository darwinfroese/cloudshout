package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

type index struct {
	Title string
}

// RenderIndex - Will be used to render our index.html web page
func RenderIndex(w http.ResponseWriter, r *http.Request) {
	i := index{Title: "Cloudshout Official Blog"}

	fmt.Println("Proto: ", r.Proto)

	if pusher, ok := w.(http.Pusher); ok {
		if err := pusher.Push("web/main.css", nil); err != nil {
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
