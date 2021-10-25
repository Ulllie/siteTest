package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "index.html")
	})


	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request){

		login := r.FormValue("login")
		pass := r.FormValue("pass")

		fmt.Fprintf(w, "Логин: %s Пароль: %s", login, pass)
	})

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))


	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
