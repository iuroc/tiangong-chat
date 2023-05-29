package main

import (
	"apee.top/tiangong-chat/serve/route"
	"fmt"
	"net/http"
)

var httpClient = &http.Client{}

func main() {
	fs := http.FileServer(http.Dir("../static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "../static/index.html")
	})
	http.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		route.LoginRoute(w, r, httpClient)
	})
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
