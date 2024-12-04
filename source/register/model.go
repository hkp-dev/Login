package register

import "fmt"

type RegisterResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
func (r *RegisterRequest) Validate() error {
	if r.Username == "" {
		return fmt.Errorf("username is required")
	}
	if r.Email == "" {
		return fmt.Errorf("email is required")
	}
	if r.Password == "" {
		return fmt.Errorf("password is required")
	}
	return nil
}

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}
