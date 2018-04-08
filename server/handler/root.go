package handler

import (
	"fmt"
	"html"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ImageSamples(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "image samples, %q", html.EscapeString(r.URL.Path))
}

