package structs

type User struct {
	Registration UserSignIn `json:"user"`
}

type UserSignIn struct {
	ID       string `json:"id"`
	Inactive bool   `json:"inactive"`
	JwtToken string `json:"jwtToken"`
}

type UserSignUp struct {
	ID          string `json:"id"`
	Uuid        string `json:"uuid"`
	Email       string `json:"email"`
	Inactive    bool   `json:"inactive"`
	Deactivated bool   `json:"userDeactivated"`
	Dormant     bool   `json:"dormant"`
	AvatarUrl   bool   `json:"avatarUrl"`
	JwtToken    string `json:"jwtToken"`
}
