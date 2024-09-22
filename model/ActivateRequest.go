package model

type ActivateRequest struct {
	Token string `json:"subscriptionToken"`
}
