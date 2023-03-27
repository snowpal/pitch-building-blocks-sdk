package recipes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/endpoints/attributes"

	recipes2 "development/go/recipes/lib/helpers/recipes"
	log "github.com/sirupsen/logrus"

	"fmt"
)

// GetResourceAttributes sign in, get resource attributes
func GetResourceAttributes() {
	log.Info("Objective: Get resource attributes")
	_, err := recipes2.ValidateDependencies()
	if err != nil {
		return
	}

	log.Info("Sign in user, email: ", lib.DefaultEmail)
	user, err := recipes2.SignIn(lib.DefaultEmail, lib.Password)

	log.Info(".get resource attributes")
	resourceAttrs, _ := attributes.GetResourceAttrs(user.JwtToken)
	if err != nil {
		return
	}

	fmt.Println(resourceAttrs)
}
