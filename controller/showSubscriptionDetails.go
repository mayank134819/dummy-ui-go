// controller/showSubscriptionDetails.go
package controller

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)


type ShowSubscriptionDetails struct {
	logger *log.Logger
}

func NewShowSubscriptionDetails(logger *log.Logger) *ShowSubscriptionDetails {
	return &ShowSubscriptionDetails{logger: logger}
}

func (c *ShowSubscriptionDetails) Show(w http.ResponseWriter, r *http.Request) {
	// Fetch subscription data from the API
	resp, err := http.Get("http://138.3.95.230:443/20180828/subscriptions/ocid1.notreviewedplaceholder.dev.dev.amaaaaaapi24rzaax7bhk6arz3jslujgylqcf3lnrsgildtphygl3tjalloq")
	if err != nil {
		http.Error(w, "Failed to fetch subscription details", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode JSON response
	var subscriptionData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&subscriptionData); err != nil {
		http.Error(w, "Error decoding subscription data", http.StatusInternalServerError)
		return
	}

	// Render template with data
	tmpl, err := template.ParseFiles("templates/showSubscriptionDetails.html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, subscriptionData)
}

func (c *ShowSubscriptionDetails) Activate(w http.ResponseWriter, r *http.Request) {
	var reqBody map[string]string
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	jsonBody, _ := json.Marshal(map[string]string{"token": reqBody["token"]})
	resp, err := http.Post("http://138.3.95.230:443/20180828/activateSubscription", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil || resp.StatusCode != http.StatusAccepted {
		http.Error(w, "Failed to activate subscription", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Subscription activated successfully"}`))
}

