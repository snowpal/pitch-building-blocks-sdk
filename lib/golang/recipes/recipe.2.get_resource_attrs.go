package main

import (
	"development/go/recipes/lib/golang"
	attributes "development/go/recipes/lib/golang/endpoints/attributes.1"
	registration "development/go/recipes/lib/golang/endpoints/registration.1"
	log "github.com/sirupsen/logrus"

	"fmt"
)

func main() {
	log.Info(".sign in user email: ", golang.Email)
	userSignIn, err := registration.SignIn(golang.Email)
	if err != nil {
		fmt.Println(err)
	}

	log.Info(".get all keys: ", userSignIn.Registration.ID)
	resourceAttrs, _ := attributes.GetResourceAttrs(userSignIn.Registration.JwtToken)
	if err != nil {
		return
	}

	fmt.Println(resourceAttrs)
}
