package login

import "fmt"

var Users = []User{}

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"passwordhash"`
	Email        string `json:"email"`
	Is2FAEnable  bool   `json:"is2faenable"`
	SecretKey    string `json:"secretkey"`
}
type LoginRequest struct {
	Username string
	Password string
}

func (r *LoginRequest) Validate() error {
	if r.Username == "" {
		return fmt.Errorf("username is required")
	}
	if r.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

type LoginResponse struct {
	Username string
	Email    string
	Is2FA    bool
}

type LoginResponse2FA struct {
	Username string
	Email    string
	Is2FA    bool
	Secret   string
}
