package main

import (
	"development/go/recipes/lib/golang"
	attributes "development/go/recipes/lib/golang/endpoints/attributes.1"
	registration "development/go/recipes/lib/golang/endpoints/registration.1"
	"development/go/recipes/lib/golang/structs/request"
	log "github.com/sirupsen/logrus"

	"fmt"
)

// sign in, get resource attributes
func main() {
	log.Info("PURPOSE: ")
	log.Info(".sign in user email: ", Email)
	reqBody := request.SignInReqBody{
		Email:    Email,
		Password: golang.Password,
	}
	userSignIn, err := registration.SignInByEmail(reqBody)
	if err != nil {
		fmt.Println(err)
	}

	log.Info(".get resource attributes")
	resourceAttrs, _ := attributes.GetResourceAttrs(userSignIn.JwtToken)
	if err != nil {
		return
	}

	fmt.Println(resourceAttrs)
}
