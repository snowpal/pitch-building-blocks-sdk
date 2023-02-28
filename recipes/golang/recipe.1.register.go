package main

import (
	keys "development/go/recipes/endpoints/keys.1"
	registration "development/go/recipes/endpoints/registration.1"
	"development/go/recipes/structs"
	"fmt"
	log "github.com/sirupsen/logrus"
)

const Email = "api_krish_4@yopmail.com"

func main() {
	userSignUp, err := registration.Signup(Email)
	if err != nil {
		fmt.Println(err)
	}

	log.Info("user ID: %s", userSignUp.ID)
	registration.Activate(userSignUp.ID)
	var userSignIn structs.User
	userSignIn, err = registration.SignIn(Email)
	if err != nil {
		fmt.Println(err)
	}

	log.Info("userRegistration: %s", userSignIn.Registration.ID)
	keys.GetAllKeys(userSignIn.Registration.JwtToken)
}
