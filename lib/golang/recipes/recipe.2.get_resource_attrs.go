package main

import (
	"development/go/recipes/lib/golang"
	attributes "development/go/recipes/lib/golang/endpoints/attributes.1"
	"development/go/recipes/lib/golang/helpers/recipes"
	log "github.com/sirupsen/logrus"

	"fmt"
)

// sign in, get resource attributes
func main() {
	log.Info("Objective: Get resource attributes")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}
	
	log.Info("Sign in user, email: ", golang.DefaultEmail)
	user, err := recipes.SignIn(golang.DefaultEmail, golang.Password)

	log.Info(".get resource attributes")
	resourceAttrs, _ := attributes.GetResourceAttrs(user.JwtToken)
	if err != nil {
		return
	}

	fmt.Println(resourceAttrs)
}
