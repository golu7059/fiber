package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	clientID         = "client-id-123"
	clientSecret     = "client-secret-abc"
	callbackCodeURL  = "http://localhost:8080/code"
	callbackTokenURL = "http://localhost:8080/token"
)

func main() {
	http.HandleFunc("/", handleLoginPage)
	http.HandleFunc("/login", handleAuthServer)
	http.HandleFunc("/code", handleCodeCallback)
	http.HandleFunc("/token", handleTokenCallback)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleLoginPage(w http.ResponseWriter, r *http.Request) {
	loginHTML := `<a href="/login">Login with OAuth</a>`
	w.Write([]byte(loginHTML))
}

func handleAuthServer(w http.ResponseWriter, r *http.Request) {
	authCode := "auth-code-1234"
	redirectURL := fmt.Sprintf("%s?code=%s", callbackCodeURL, authCode)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func handleCodeCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	if code != "auth-code-1234" {
		http.Error(w, "Invalid authorization code", http.StatusUnauthorized)
		return
	}

	accessToken := "access-token-xyz"
	redirectURL := fmt.Sprintf("%s?token=%s", callbackTokenURL, accessToken)
	http.Redirect(w, r, redirectURL, http.StatusFound)
}

func handleTokenCallback(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "Token not found", http.StatusUnauthorized)
		return
	}
	w.Write([]byte(fmt.Sprintf("Access Token: %s", token)))
}