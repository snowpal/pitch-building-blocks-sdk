package main

import (
	"development/go/recipes/lib/building-blocks"
	attributes "development/go/recipes/lib/building-blocks/endpoints/attributes.1"
	"development/go/recipes/lib/building-blocks/helpers/recipes"
	log "github.com/sirupsen/logrus"

	"fmt"
)

// GetResourceAttributes sign in, get resource attributes
func GetResourceAttributes() {
	log.Info("Objective: Get resource attributes")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	log.Info("Sign in user, email: ", building_blocks.DefaultEmail)
	user, err := recipes.SignIn(building_blocks.DefaultEmail, building_blocks.Password)

	log.Info(".get resource attributes")
	resourceAttrs, _ := attributes.GetResourceAttrs(user.JwtToken)
	if err != nil {
		return
	}

	fmt.Println(resourceAttrs)
}
