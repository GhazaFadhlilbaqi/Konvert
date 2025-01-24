package handlers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Konvert v0.1")
}

func Converter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Paste a YT link here")
}
