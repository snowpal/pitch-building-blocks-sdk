package structs

type UserRegistration struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	JwtToken string `json:"jwtToken"`
}
