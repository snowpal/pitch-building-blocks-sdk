package response

type Users struct {
	Users []User `json:"users"`
}

type UserRegistration struct {
	User User `json:"user"`
}

type User struct {
	ID          string `json:"id"`
	Uuid        string `json:"uuid"`
	Email       string `json:"email"`
	Inactive    bool   `json:"inactive"`
	Deactivated bool   `json:"userDeactivated"`
	Dormant     bool   `json:"dormant"`
	AvatarUrl   string `json:"avatarUrl"`
	JwtToken    string `json:"jwtToken"`
}

type SearchUsers struct {
	SearchUsers []SearchUser `json:"users"`
}

type SearchUser struct {
	ID        string `json:"id"`
	ProfileID string `json:"profileId"`
	Username  string `json:"userName"`
	FirstName string `json:"firstName"`
	FullName  string `json:"fullName"`
}
