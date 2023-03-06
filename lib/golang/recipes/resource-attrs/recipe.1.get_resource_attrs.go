package main

import (
	"development/go/recipes/lib/golang"
	attributes "development/go/recipes/lib/golang/endpoints/attributes.1"
	registration "development/go/recipes/lib/golang/endpoints/registration.1"
	log "github.com/sirupsen/logrus"

	"fmt"
)

// sign in, get resource attributes
func main() {
	log.Info(".sign in user email: ", golang.Email)
	userSignIn, err := registration.SignIn(golang.Email)
	if err != nil {
		fmt.Println(err)
	}

	log.Info(".get resource attributes")
	resourceAttrs, _ := attributes.GetResourceAttrs(userSignIn.User.JwtToken)
	if err != nil {
		return
	}

	fmt.Println(resourceAttrs)
}
