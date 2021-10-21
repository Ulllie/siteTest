package main

import (
	"fmt"
	"net/http"
)

func home_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Привет pidor!")
}

func main() {
	http.HandleFunc("/", home_page)
	http.ListenAndServe(":7070", nil)
}
