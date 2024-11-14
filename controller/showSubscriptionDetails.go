// controller/showSubscriptionDetails.go
package controller

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)


type ShowSubscriptionDetails struct {
	logger *log.Logger
}

func NewShowSubscriptionDetails(logger *log.Logger) *ShowSubscriptionDetails {
	return &ShowSubscriptionDetails{logger: logger}
}

func (c *ShowSubscriptionDetails) Show(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	subscriptionToken := vars["subscriptionToken"]

	c.logger.Println("Received request for ShowSubscriptionDetails")
	c.logger.Println("Subscription Token:", subscriptionToken)

	var apiURL string
	if subscriptionToken != "" {
		// Use the subscriptionToken in the API URL
		apiURL = fmt.Sprintf("http://138.3.95.230:443/20180828/subscriptions/%s", subscriptionToken)
		c.logger.Println("API URL with subscriptionToken:", apiURL)
	} else {
		// Default URL when no subscriptionToken is provided
		apiURL = "http://138.3.95.230:443/20180828/subscriptions/ocid1.notreviewedplaceholder.dev.dev.amaaaaaapi24rzaax7bhk6arz3jslujgylqcf3lnrsgildtphygl3tjalloq"
		c.logger.Println("API URL with subscriptionToken:", apiURL)
	}
	apiURL = "http://138.3.95.230:443/20180828/subscriptions/ocid1.notreviewedplaceholder.dev.dev.amaaaaaapi24rzaax7bhk6arz3jslujgylqcf3lnrsgildtphygl3tjalloq"



	// Fetch subscription data from the API
	// Fetch subscription data from the API
	c.logger.Println("Making GET request to:", apiURL)
	// resp, err := http.Get("http://138.3.95.230:443/20180828/subscriptions/ocid1.notreviewedplaceholder.dev.dev.amaaaaaapi24rzaax7bhk6arz3jslujgylqcf3lnrsgildtphygl3tjalloq")
	resp, err := http.Get(apiURL)
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
	c.logger.Println("Decoded subscription data:", subscriptionData)


	// Convert subscriptionData to JSON to embed in HTML
	subscriptionDataJSON, err := json.Marshal(subscriptionData)
	if err != nil {
		c.logger.Println("Error marshaling subscription data to JSON:", err)
		http.Error(w, "Error processing subscription data", http.StatusInternalServerError)
		return
	}

	// Render template with data
	c.logger.Println("Parsing template 'showSubscriptionDetails.html'")
	tmpl, err := template.ParseFiles("templates/showSubscriptionDetails.html")
	if err != nil {
		c.logger.Println("Error parsing template:", err)
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}

	// Pass the JSON data to the template
	c.logger.Println("Rendering template with subscription data")
	err = tmpl.Execute(w, map[string]interface{}{
		"SubscriptionDataJSON": template.JS(subscriptionDataJSON), // Pass JSON as a JS-safe string
	})
	if err != nil {
		c.logger.Println("Error executing template:", err)
	}
	c.logger.Println("Template rendered successfully")



	// // Render template with data
	// tmpl, err := template.ParseFiles("templates/showSubscriptionDetails.html")
	// if err != nil {
	// 	http.Error(w, "Template not found", http.StatusInternalServerError)
	// 	return
	// }
	// tmpl.Execute(w, subscriptionData)
}

// func (c *ShowSubscriptionDetails) Activate(w http.ResponseWriter, r *http.Request) {
// 	var reqBody map[string]string
// 	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
// 		http.Error(w, "Invalid request", http.StatusBadRequest)
// 		return
// 	}

// 	jsonBody, _ := json.Marshal(map[string]string{"token": reqBody["token"]})
// 	resp, err := http.Post("http://138.3.95.230:443/20180828/activateSubscription", "application/json", bytes.NewBuffer(jsonBody))
// 	if err != nil || resp.StatusCode != http.StatusAccepted {
// 		http.Error(w, "Failed to activate subscription", http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(`{"message": "Subscription activated successfully"}`))
// }

func (c *ShowSubscriptionDetails) Activate(w http.ResponseWriter, r *http.Request) {
    var reqBody map[string]string
    if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
        log.Printf("Error decoding request body: %v", err)
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }

    // Log the received selfTokenId
    selfTokenId, ok := reqBody["selfTokenId"]
    if !ok {
        log.Println("selfTokenId not provided in request body")
        http.Error(w, "selfTokenId is required", http.StatusBadRequest)
        return
    }
    log.Printf("Activating subscription with selfTokenId: %s", selfTokenId)

    // Prepare the JSON body with the selfTokenId
    jsonBody, _ := json.Marshal(map[string]string{"selfTokenId": selfTokenId})

    // Log the request URL and payload
    url := "http://138.3.95.230:443/20180828/partner/subscriptions/actions/activate"
    log.Printf("Sending PUT request to URL: %s with body: %s", url, jsonBody)

    // Perform the PUT request to the specified endpoint
    client := &http.Client{}
    req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonBody))
    if err != nil {
        log.Printf("Error creating PUT request: %v", err)
        http.Error(w, "Failed to create request", http.StatusInternalServerError)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    // Send the request
    resp, err := client.Do(req)
    if err != nil {
        log.Printf("Error sending PUT request: %v", err)
        http.Error(w, "Failed to activate subscription", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    // Log the response status
    log.Printf("Received response with status code: %d", resp.StatusCode)
    if resp.StatusCode != http.StatusAccepted {
        log.Printf("Failed to activate subscription, status code: %d", resp.StatusCode)
        http.Error(w, "Failed to activate subscription", http.StatusInternalServerError)
        return
    }

    // Successful activation log
    log.Println("Subscription activated successfully")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "Subscription activated successfully"}`))
}
