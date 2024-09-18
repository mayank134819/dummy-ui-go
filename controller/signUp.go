package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"oracle.com/self/partner-dummy-env/model"
)

type SignUp struct {
	logger *log.Logger
}

func NewSingUp(logger *log.Logger) *SignUp {
	return &SignUp{
		logger: logger,
	}
}

func (su *SignUp) SignUp(w http.ResponseWriter, r *http.Request) {
	su.logger.Println("Sing up redirect")
	// Make the POST request to activate
	requestBody := &model.ActivateRequest{
		Token: "abcdef",
	}
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		su.logger.Println("Error marshaling JSON:", err)
		return
	}
	resp, err := http.Post("https://example.com/api/users", "application/json", bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		su.logger.Println("Error making POST request:", err)
		// return
	}
	defer resp.Body.Close()

	// Redirect on SELF Page
	http.Redirect(w, r, "https://bitbucket.oci.oraclecorp.com/users/rohitprs/repos/partner-test-env/browse", http.StatusTemporaryRedirect)
}
