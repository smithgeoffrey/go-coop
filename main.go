package main

import (
	"net/http"
)

func main() {

	// Handlers for pages
	http.HandleFunc("/", index)
	http.HandleFunc("/signin", signin)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/other", other)
	http.HandleFunc("/favicon.ico", favicon)

	// Handler for file-serving the `public` dir.  Prefix stripping is
	// needed to be able to do this in our templates:
	//
	//    <img src="/public/images/image1.jpg">
	//
	dir := http.Dir("public")
	fileServer := http.FileServer(dir)
	handlerWithStrippedPrefix := http.StripPrefix("/public/", fileServer)
	http.Handle("/public/", handlerWithStrippedPrefix)

	// Start server with the DefaultServeMux
	http.ListenAndServe(":8080", nil)

}
