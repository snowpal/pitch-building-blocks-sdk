package response

type UserSignedIn struct {
	Registration UserSignIn `json:"user"`
}

type UserSignedUp struct {
	Registration UserSignUp `json:"user"`
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
	AvatarUrl   string `json:"avatarUrl"`
	JwtToken    string `json:"jwtToken"`
}

type User struct {
	ID        string `json:"userId"`
	ProfileID string `json:"profileId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Initial   string `json:"userInitial"`
	Email     string `json:"email"`
}
