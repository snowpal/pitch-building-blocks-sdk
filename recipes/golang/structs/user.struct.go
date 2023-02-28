package structs

type User struct {
	Registration UserSignIn `json:"user"`
}

type UserSignIn struct {
	ID       string `json:"id"`
	Inactive bool   `json:"inactive"`
	JwtToken string `json:"jwtToken"`
}
