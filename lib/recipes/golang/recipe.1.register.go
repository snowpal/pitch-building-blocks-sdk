package main

import (
	keys "development/go/recipes/endpoints/keys.1"
	registration "development/go/recipes/endpoints/registration.1"
	log "github.com/sirupsen/logrus"

	"development/go/recipes/structs"
	"fmt"
)

const Email = "api_krish_16@yopmail.com"

func main() {
	log.Info(".sign up user with email: ", Email)
	userSignUp, err := registration.Signup(Email)
	if err != nil {
		fmt.Println(err)
	}

	log.Info(".activate user ID: ", userSignUp.Registration.ID)
	registration.Activate(userSignUp.Registration.ID)

	log.Info(".sign in user ID: ", userSignUp.Registration.ID)
	var userSignIn structs.UserSignedIn
	userSignIn, err = registration.SignIn(Email)
	if err != nil {
		fmt.Println(err)
	}

	log.Info(".get all keys: ", userSignIn.Registration.ID)
	keys.GetAllKeys(userSignIn.Registration.JwtToken)
}
