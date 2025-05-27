package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// LoginRequest represents the JSON payload from the login form
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// Serve static files
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Handle login requests
	http.HandleFunc("/login", handleLogin)

	// Start the server
	fmt.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Only allow POST requests
	if r.Method != http.MethodPost {
		fmt.Println("Method not allowed:", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Log request headers for debugging
	fmt.Println("Content-Type:", r.Header.Get("Content-Type"))
	
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	
	// Log the raw request body for debugging
	fmt.Println("Request body:", string(body))
	
	// Parse the JSON request
	var loginReq LoginRequest
	err = json.Unmarshal(body, &loginReq)
	if err != nil {
		fmt.Println("JSON parse error:", err)
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	fmt.Println("Parsed username:", loginReq.Username)
	fmt.Println("Parsed password:", loginReq.Password)

	// Simple validation (in a real app, you'd check against a database)
	if strings.TrimSpace(loginReq.Username) == "" || strings.TrimSpace(loginReq.Password) == "" {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`
			<div class="error">
				<p>Username and password cannot be empty</p>
			</div>
		`))
		return
	}

	// For demo purposes, accept any non-empty username/password
	// In a real application, you would validate against a database

	// Return an HTML response
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(fmt.Sprintf(`
		<div class="success">
			<h3>Login Successful!</h3>
			<p>Welcome, %s!</p>
			<p>You have successfully logged in.</p>
		</div>
	`, loginReq.Username)))
}
