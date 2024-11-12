package model

type SignUpRequest struct {
	Token    string `json:"subscriptionToken"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
