package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, rw *http.Request) {
	var filename = "login.html"
	t, err := template.ParseFiles(filename)
	if err != nil {
		fmt.Println("Error parsing file", err)
		return
	}
	t.ExecuteTemplate(w, filename, nil)
	if err != nil {
		fmt.Println("Error executing template", err)
		return
	}
}

var userDB = map[string]string{
	"blacknyangumi": "forgotpassword",
}

func login_submit(w http.ResponseWriter, rw *http.Request) {
	username := rw.FormValue("username")
	password := rw.FormValue("password")

	if userDB[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "You are now logged in")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "You've entered wrong username or password")
	}

}

func handler(w http.ResponseWriter, rw *http.Request) {
	switch rw.URL.Path {
	case "/login":
		login(w, rw)
	case "/login-submit":
		login_submit(w, rw)
	}

}

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Error starting server", err)
		return
	}
}
