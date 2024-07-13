package middleware

import (
	"log"
	"time"
	"crypto/rand"
	"net/http"
	"fmt"

	"github.com/o1egl/paseto"

	m "qualitech.paseto-auth/model"
)


var symmetricKey = generateKey()


func generateKey() []byte {
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}
	return key
}

func GetSymeticKey() []byte {
    return symmetricKey
}

func GenerateToken(username string) (string, error) {
	now := time.Now()
	jsonToken := m.JSONToken{
		Issuer:     "example.com",
		Subject:    username,
		Audience:   "example audience",
		IssuedAt:   now,
		Expiration: now.Add(24 * time.Hour),
	}

	footer := "some-footer"
	v2 := paseto.NewV2()
	token, err := v2.Encrypt(symmetricKey, jsonToken, footer)
	return token, err
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Authorization header required")
			return
		}

		var jsonToken m.JSONToken
		var footer string
		err := paseto.NewV2().Decrypt(tokenString, GetSymeticKey(), &jsonToken, &footer)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Invalid token")
			return
		}

		if jsonToken.Expiration.Before(time.Now()) {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Token expired")
			return
		}

		next(w, r)
	}
}