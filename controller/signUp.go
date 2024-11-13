package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"oracle.com/self/partner-test-env/model"
	"oracle.com/self/partner-test-env/database"
)

type SignUp struct {
	logger       *log.Logger
	sessionStore *sessions.CookieStore
}

func NewSignUp(logger *log.Logger, sessionStore *sessions.CookieStore) *SignUp {
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
		http.Error(w, "Invalid token md", http.StatusBadRequest)
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
	resp, err := client.Post("http://138.3.95.230:443/20180828/subscriptions/resolve", "application/json", bytes.NewBuffer(jsonRequestBody))
	if err != nil || resp.StatusCode != 202 {
		su.logger.Println("Error making POST request:", err)
		http.Error(w, "Invalid token dm", http.StatusBadRequest)
		return
	}
	su.logger.Println("Response is ", resp)
	defer resp.Body.Close()

	// Read the complete response body
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		su.logger.Println("Error reading response body:", err)
		http.Error(w, "Failed to read response", http.StatusInternalServerError)
		return
	}
	su.logger.Println("Complete Response Data:", string(responseData))

	su.logger.Println("Form data is:", reqBody.Email )
	su.logger.Println("Form data is:", reqBody.Password )
	// su.logger.Println("Token fetched is: ", responseData.subscriptionToken)


	// Fetch the user from the database
	user, err := database.GetUser(reqBody.Email, reqBody.Password)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Read and parse the response body
	var responseMap map[string]interface{}
	err = json.Unmarshal(responseData, &responseMap)
	if err != nil {
		su.logger.Println("Error decoding response body:", err)
		http.Error(w, "Failed to parse response", http.StatusInternalServerError)
		return
	}

	su.logger.Println("Complete Response Data:", responseMap)

	// Extract the selfToken from the response
	selfToken, ok := responseMap["selfTokenId"].(string)
	if !ok || selfToken == "" {
		su.logger.Println("Subscription token not found in response")
		http.Error(w, "Failed to retrieve subscription token", http.StatusInternalServerError)
		return
	}

	

	// Construct the showSubscriptionDetails URL with the subscriptionToken
	showSubscriptionDetailsURL := "/showSubscriptionDetails/" + selfToken
	su.logger.Println("Redirecting to:", showSubscriptionDetailsURL)

	// http.Redirect(w, r, showSubscriptionDetailsURL, http.StatusSeeOther)

	// Set response headers for JSON and CORS (if necessary)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Send the URL as JSON in the response
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "PendingActivation",
		"url":     showSubscriptionDetailsURL,
	})





	

	// Create a new session and store user data
	// session, err := su.sessionStore.Get(r, "session-name")
	// if err != nil {
	// 	http.Error(w, "Failed to create session", http.StatusInternalServerError)
	// 	return
	// }

	// // Store user ID and email in the session
	// session.Values["userID"] = user.ID
	// session.Values["email"] = user.Email
	// session.Save(r, w)

	su.logger.Println("User signed in successfully:", user.Email)

	///////
	
	
	



	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	// w.Header().Set("Access-Control-Allow-Origin", "*")                            // Allow all origins, or specify your domain
	// w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")          // Allow specific methods
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization") // Allow specific headers

	// // Response
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusAccepted)
	// json.NewEncoder(w).Encode(map[string]string{"message": "Activated"})
}
