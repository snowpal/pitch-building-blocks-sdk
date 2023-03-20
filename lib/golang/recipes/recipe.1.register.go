package main

import (
	"development/go/recipes/lib/golang"
	keys "development/go/recipes/lib/golang/endpoints/keys.1"
	registration "development/go/recipes/lib/golang/endpoints/registration.1"
	"development/go/recipes/lib/golang/structs/request"
	log "github.com/sirupsen/logrus"

	"fmt"
)

const Email = "apiuser3@gmail.com"

// Sign up, activate user, sign in, get all keys.
func main() {
	log.Info(".sign up user with email: ", Email)
	signUpReqBody := request.SignupReqBody{
		Email:           Email,
		Password:        golang.Password,
		ConfirmPassword: golang.Password,
	}
	user, err := registration.RegisterNewUserByEmail(signUpReqBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Info(".activate user ID: ", user.ID)
	err = registration.ActivateUser(user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Info(".sign in user ID: ", user.ID)
	signInReqBody := request.SignInReqBody{
		Email:    Email,
		Password: golang.Password,
	}

	user, err = registration.SignInByEmail(signInReqBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Info(".get all keys: ", user.ID)
	allKeys, err := keys.GetKeys(user.JwtToken, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Info(allKeys)
}
