package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, send a message to " +
		"<a href=\"mailto:support@lenslocked.com\">" +
		"support@lenslocked.com</a>.")
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "This page answers the most frequently " +
    "asked questions about LensLocked.")
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>We couldn't find the page you're " +
	" looking for.</h1>" + "<p>Please " +
	"<a href=\"mailto:support@lenslocked.com\">email us</a> if " + 
	"you continue to experience this issue.</p>")
}

func main() {
	var notFound http.Handler = http.HandlerFunc(pageNotFound)
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.NotFoundHandler = notFound
	http.ListenAndServe(":3000", r)
}