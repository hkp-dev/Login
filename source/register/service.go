package register

import (
	"ServerAPI/source/login"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/rand"

	"github.com/pquerna/otp/totp"
)

func AddUserWithSecret(username, password, email string) (string, string, error) {
	for _, user := range login.Users {
		if user.Username == username {
			return "", "", errors.New("Invalid username")
		}
	}

	hashedPassword := hashPassword(password)

	secretKey, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "App",
		AccountName: email,
	})
	if err != nil {
		return "", "", fmt.Errorf("failed to generate secret key: %v", err)
	}

	qrCodePath := fmt.Sprintf("./views/templates/qrcodes/%d.png", rand.Intn(100000))

	newUser := login.User{
		ID:           len(login.Users) + 1,
		Username:     username,
		PasswordHash: hashedPassword,
		Email:        email,
		Is2FAEnable:  false,
		SecretKey:    secretKey.Secret(),
	}
	login.Users = append(login.Users, newUser)

	return secretKey.Secret(), qrCodePath, nil
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}
