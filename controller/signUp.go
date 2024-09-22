package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"oracle.com/self/partner-dummy-env/model"
)

type SignUp struct {
	logger       *log.Logger
	sessionStore *sessions.CookieStore
}

func NewSingUp(logger *log.Logger, sessionStore *sessions.CookieStore) *SignUp {
	return &SignUp{
		logger:       logger,
		sessionStore: sessionStore,
	}
}

func (su *SignUp) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	var reqBody model.SignUpRequest
	// Decode the JSON body
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	su.logger.Println("Form request value =", reqBody)

	requestBody := &model.ActivateRequest{}

	if reqBody.Token == "" {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	// Create Request body for activation call.
	requestBody.Token = reqBody.Token
	su.logger.Println("Self requestBody token is = ", requestBody.Token)
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		su.logger.Println("Error marshaling JSON:", err)
		return
	}
	su.logger.Println("jsonRequestBody = ", string(jsonRequestBody))

	// Create a new HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout of 5 seconds
	}
	resp, err := client.Post("http://138.3.85.250:443/20180828/subscriptions/resolve", "application/json", bytes.NewBuffer(jsonRequestBody))
	if err != nil || resp.StatusCode != 202 {
		su.logger.Println("Error making POST request:", err)
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}
	su.logger.Println(resp)
	defer resp.Body.Close()

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")                            // Allow all origins, or specify your domain
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allow specific methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow specific headers

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{"message": "Activated"})
}
