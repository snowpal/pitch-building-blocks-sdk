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

	log.Info(".TODO(): get a key, and use it below")
	attributes.UpdateKeyAttrs(userSignIn.Registration.JwtToken, "")

	log.Info(".TODO(): get a block, and use it below")
	attributes.UpdateBlockAttrs(userSignIn.Registration.JwtToken, "")

	log.Info(".TODO(): get a pod, and use it below")
	attributes.UpdatePodAttrs(userSignIn.Registration.JwtToken, "")
}
