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

	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	su.logger.Println("username = ", r.FormValue("username"), "password = ", r.FormValue("password"))

	session, _ := su.sessionStore.Get(r, "selfReg")
	selfToken := session.Values["selfToken"]
	su.logger.Println("Get token value from session = ", selfToken)

	requestBody := &model.ActivateRequest{}

	if str, ok := selfToken.(string); ok && str != "" {
		requestBody.Token = str
	} else {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	su.logger.Println("Self requestBody token is = ", requestBody.Token)
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		su.logger.Println("Error marshaling JSON:", err)
		return
	}
	// Create a new HTTP client with a timeout
	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout of 5 seconds
	}
	resp, err := client.Post("http://138.3.72.134:443/20180828/subscriptions/resolve", "application/json", bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		su.logger.Println("Error making POST request:", err, jsonRequestBody)
		// return
	}
	su.logger.Println(resp)
	defer resp.Body.Close()

	// Redirect on SELF Page
	http.Redirect(w, r, "https://oracle.com", http.StatusTemporaryRedirect)
}
