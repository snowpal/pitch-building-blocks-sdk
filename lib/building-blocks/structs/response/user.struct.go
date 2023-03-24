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

type TaggedUser struct {
	ID        string `json:"userId"`
	ProfileID string `json:"profileId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Initial   string `json:"userInitial"`
	Email     string `json:"email"`
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

type SharedUser struct {
	ID            string `json:"id"`
	Acl           string `json:"acl"`
	Username      string `json:"userName"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	FullName      string `json:"fullName"`
	AddressUserBy string `json:"addressUserBy"`
}
