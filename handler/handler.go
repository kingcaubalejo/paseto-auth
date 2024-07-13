package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"

	m "qualitech.paseto-auth/model"
	wrapper "qualitech.paseto-auth/middleware"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	var loginReq m.LoginRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error reading request body")
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &loginReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid request body")
		return
	}

	// Dummy user validation
	if loginReq.Username == "user" && loginReq.Password == "password" {
		token, err := wrapper.GenerateToken(loginReq.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Could not generate token")
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", token)
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "Invalid credentials")
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are authenticated")
}

func UnProtectedHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Open Api")
}