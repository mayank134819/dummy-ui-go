package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"oracle.com/self/partner-test-env/database"
	"oracle.com/self/partner-test-env/model"
	"github.com/gorilla/sessions"

)

type SignIn struct {
	logger       *log.Logger
	sessionStore *sessions.CookieStore
}

func NewSignIn(logger *log.Logger, sessionStore *sessions.CookieStore) *SignIn {
	return &SignIn{
		logger:       logger,
		sessionStore: sessionStore,
	}
}

func (si *SignIn) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	si.logger.Println("Coming here ************")

	// Parse form values from the sign-up form
	// name := r.FormValue("name")
	// email := r.FormValue("email")
	// password := r.FormValue("password")

	// Decode the JSON request body into the SignUpRequest struct
	// var req SignUpRequest
	var req model.SignUpRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		si.logger.Println("Failed to decode JSON body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}



	// Fetch the 'st' query parameter from the URL
	// token := r.URL.Query().Get("st")
	si.logger.Println("///////////////details :////////////////", req.Email)
	si.logger.Println("///////////////fetched token :////////////////", req.Token)
	// if token == "" {
	// 	http.Error(w, "Missing token in URL", http.StatusBadRequest)
	// 	return
	// }

	// Insert the user into the database
	err = database.CreateUser(req.Name, req.Email, req.Password)
	if err != nil {
		si.logger.Println("///////////////Error creating user:////////////////", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}


	si.logger.Println("*************Userrrrrrrr is created **************:", err)
	// si.logger.Println("Fetched token from URL:", token)

	redirectURL := "/?st=" + req.Token
	si.logger.Println("Fetched token from URL:", redirectURL)
	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
