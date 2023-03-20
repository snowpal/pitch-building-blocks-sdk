package request

type SignupReqBody struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type SignInReqBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResetPasswordReqBody struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}
