package request

type UserRegistrationReqBody struct {
	Email           *string `json:"email"`
	Password        string  `json:"password"`
	ConfirmPassword *string `json:"confirmPassword"`
}
