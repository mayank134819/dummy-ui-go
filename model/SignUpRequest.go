package model

type SignUpRequest struct {
	Token    string `json:"subscriptionToken"`
	Username string `json:"username"`
	Password string `json:"password"`
}
