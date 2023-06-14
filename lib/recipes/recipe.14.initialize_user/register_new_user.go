package recipes

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GenerateDynamicEmail() string {
	var email string
	return email
}

func RegisterNewUser() (string, error) {
	var err error
	var user response.User
	for registered := true; registered; registered = err != nil {
		user, err = recipes.RegisterUser(GenerateDynamicEmail())
	}
	return user.Email, nil
}
