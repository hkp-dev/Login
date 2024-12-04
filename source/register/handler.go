package register

import (
	"fmt"
	"html/template"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("./views/templates/register.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		repeatPassword := r.FormValue("repeatPassword")
		email := r.FormValue("mail")

		if password != repeatPassword {
			http.Error(w, "Passwords do not match", http.StatusBadRequest)
			return
		}

		secretKey, qrCode, err := AddUserWithSecret(username, password, email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Println("SecretKey: ", secretKey)
		data := map[string]interface{}{
			"QRCodePath": qrCode,
			// "SecretKey":  secretKey,
		}
		tmpl, err := template.ParseFiles("./views/templates/qr.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}
