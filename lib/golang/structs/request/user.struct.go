package request

type UserRegistration struct {
	Email           *string `json:"email"`
	Password        string  `json:"password"`
	ConfirmPassword *string `json:"confirmPassword"`
}
