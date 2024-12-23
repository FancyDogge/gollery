package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>ZULUL</h1> <p>Username: %s</p>", username)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w,
		`<h1>FAQ Page</h1>
		<ul>
			<li>Q: blah blah blah?</li>
			<li>A: bladfbldjbladlfb</li>
			<li>Q: blah blah blah?</li>
			<li>A: bladfbldjbladlfb</li>
			<li>Q: blah blah blah?</li>
			<li>A: bladfbldjbladlfb</li>
		</ul>
		`)
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	case "/faq":
// 		faqHandler(w, r)
// 	default:
// 		http.Error(w, "Sorry, page not found!", http.StatusNotFound)
// 	}
// }

func main() {
	//вместо кастомного роутера на коленке, пока просто как пример
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/{username}", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Sorry, page not found!", http.StatusNotFound)
	})
	fmt.Println("Starting the server on port :3000...")
	http.ListenAndServe(":3000", r)
}
