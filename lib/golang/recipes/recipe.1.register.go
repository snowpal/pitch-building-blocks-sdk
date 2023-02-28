package main

import (
	"development/go/recipes/lib/golang"
	keys "development/go/recipes/lib/golang/endpoints/keys.1"
	registration "development/go/recipes/lib/golang/endpoints/registration.1"
	"development/go/recipes/lib/golang/structs"
	log "github.com/sirupsen/logrus"

	"fmt"
)

func main() {
	log.Info(".sign up user with email: ", golang.Email)
	userSignUp, err := registration.Signup(golang.Email)
	if err != nil {
		fmt.Println(err)
	}

	log.Info(".activate user ID: ", userSignUp.Registration.ID)
	registration.Activate(userSignUp.Registration.ID)

	log.Info(".sign in user ID: ", userSignUp.Registration.ID)
	var userSignIn structs.UserSignedIn
	userSignIn, err = registration.SignIn(golang.Email)
	if err != nil {
		fmt.Println(err)
	}

	log.Info(".get all keys: ", userSignIn.Registration.ID)
	keys.GetAllKeys(userSignIn.Registration.JwtToken)
}
