package main

import (
	keys "development/go/recipes/lib/golang/endpoints/keys.1"
	"development/go/recipes/lib/golang/endpoints/registration.1"
	"development/go/recipes/lib/golang/endpoints/structs"
	log "github.com/sirupsen/logrus"

	"fmt"
)

const Email = "api_krish_19@yopmail.com"

func main() {
	log.Info(".sign up user with email: ", Email)
	userSignUp, err := registration_1.Signup(Email)
	if err != nil {
		fmt.Println(err)
	}

	log.Info(".activate user ID: ", userSignUp.Registration.ID)
	registration_1.Activate(userSignUp.Registration.ID)

	log.Info(".sign in user ID: ", userSignUp.Registration.ID)
	var userSignIn structs.UserSignedIn
	userSignIn, err = registration_1.SignIn(Email)
	if err != nil {
		fmt.Println(err)
	}

	log.Info(".get all keys: ", userSignIn.Registration.ID)
	keys.GetAllKeys(userSignIn.Registration.JwtToken)
}
