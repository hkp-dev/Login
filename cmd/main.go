package main

import (
	"ServerAPI/source/login"
	"ServerAPI/source/register"
	"log"
	"net/http"
)

// var (
// 	MethodGet  = "GET"
// 	MethodPost = "POST"
// 	MethodPut  = "PUT"
// 	Users      = []User{}
// )

// type User struct {
// 	ID           int
// 	Username     string
// 	PasswordHash string
// 	Email        string
// 	Is2FAEnable  bool
// 	SecretKey    string
// }

func main() {
	router := http.NewServeMux()
	// router.HandleFunc("/docs", docsHandler)
	// router.HandleFunc("/health", healthHandler)
	router.HandleFunc("/", login.LoginHandler)
	router.HandleFunc("/register", register.RegisterHandler)
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

// func registerHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == MethodGet {
// 		tmpl, err := template.ParseFiles("./views/templates/register.html")
// 		if err != nil {
// 			fmt.Printf("Error parsing template in /register: %v\n", err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		if err := tmpl.Execute(w, nil); err != nil {
// 			fmt.Printf("Error executing template in /register: %v\n", err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "text/html")
// 	}
// 	if r.Method == MethodPost {
// 		username := r.FormValue("username")
// 		password := r.FormValue("password")
// 		repeatPassword := r.FormValue("repeatPassword")
// 		email := r.FormValue("mail")
// 		if password != repeatPassword {
// 			w.WriteHeader(http.StatusBadRequest)
// 			fmt.Fprintln(w, "Passwords do not match")
// 			return
// 		}
// 		hashedPassword := hashPassword(password)
// 		userID := len(users) + 1
// 		newUser := User{
// 			ID:           userID,
// 			Username:     username,
// 			PasswordHash: hashedPassword,
// 			Email:        email,
// 			Is2FAEnable:  false,
// 		}
// 		users = append(users, newUser)
// 		fmt.Printf("User registered: %v\n", newUser)
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintln(w, "User registered successfully")
// 		return
// 	}
// }
// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == MethodGet {
// 		tmpl, err := template.ParseFiles("./views/templates/login.html")
// 		if err != nil {
// 			fmt.Printf("Error parsing template in /login: %v\n", err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		if err := tmpl.Execute(w, nil); err != nil {
// 			fmt.Printf("Error executing template in /login: %v\n", err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "text/html")
// 	}
// 	if r.Method == MethodPost {
// 		username := r.FormValue("username")
// 		password := r.FormValue("password")
// 		fmt.Println(username)
// 		fmt.Println(password)
// 		w.WriteHeader(http.StatusOK)
// 		if username == "" || password == "" {
// 			w.WriteHeader(http.StatusBadRequest)
// 			fmt.Fprintln(w, "Username or password is empty")
// 			return
// 		}
// 		if username == "admin" && password == "admin" {
// 			w.WriteHeader(http.StatusOK)
// 			fmt.Fprintln(w, "Login success")
// 			return
// 		}
// 	}
// }
// func healthHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	if w == nil {
// 		return
// 	}
// 	if _, err := w.Write([]byte("OK")); err != nil {
// 		panic(err)
// 	}
// }
// func docsHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != MethodGet && r.Method != MethodPut && r.Method != http.MethodPatch {
// 		w.WriteHeader(http.StatusMethodNotAllowed)
// 		fmt.Fprintln(w, "Method Not Allowed")
// 		return
// 	}

// 	tmpl, err := template.ParseFiles("../views/templates/index.html")
// 	if err != nil {
// 		fmt.Printf("Error parsing template in /docs: %v\n", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	if err := tmpl.Execute(w, nil); err != nil {
// 		fmt.Printf("Error executing template in /docs: %v\n", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "text/html")
// }
// func hashPassword(password string) string {
// 	hash := sha256.Sum256([]byte(password))
// 	return hex.EncodeToString(hash[:])
// }
