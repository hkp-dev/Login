package login

import (
	"fmt"
	"net/http"
	"text/template"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./views/templates/login.html")
		if err != nil {
			fmt.Printf("Error parsing template in /login: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			fmt.Printf("Error executing template in /login: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
	}
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		fmt.Println("Username:", username)
		fmt.Println("Password:", password)
		w.WriteHeader(http.StatusOK)
		if username == "" || password == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "Username or password is empty")
			return
		}
		if username == "admin" && password == "admin" {
			fmt.Fprintln(w, "Login success")
			return
		}
	}
}