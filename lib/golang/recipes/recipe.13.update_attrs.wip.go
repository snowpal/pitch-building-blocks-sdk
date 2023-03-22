package main

import (
	"development/go/recipes/lib/golang"
	attributes "development/go/recipes/lib/golang/endpoints/attributes.1"
	registration "development/go/recipes/lib/golang/endpoints/registration.1"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/request"
	log "github.com/sirupsen/logrus"

	"fmt"
)

// sign in, update key attributes, update block attributes, update pod attributes, update block pod attributes,
// get resource attributes
func main() {
	log.Info("Objective: ")
	log.Info(".sign in user email: ", golang.DefaultEmail)
	userSignIn, err := registration.SignInByEmail(request.SignInReqBody{
		Email:    golang.DefaultEmail,
		Password: golang.Password,
	})
	if err != nil {
		fmt.Println(err)
	}

	log.Info(".TODO(): get a key, and use it below")
	err = attributes.UpdateKeyAttrs(userSignIn.JwtToken, "", request.ResourceAttribute{})
	if err != nil {
		return
	}

	log.Info(".TODO(): get a block, and use it below")
	err = attributes.UpdateBlockAttrs(userSignIn.JwtToken, common.ResourceIdParam{}, request.ResourceAttribute{})
	if err != nil {
		return
	}

	log.Info(".TODO(): get a pod, and use it below")
	err = attributes.UpdatePodAttrs(userSignIn.JwtToken, common.ResourceIdParam{}, request.ResourceAttribute{})
	if err != nil {
		return
	}

	log.Info(".TODO(): get a block pod, and use it below")
	err = attributes.UpdateBlockPodAttrs(userSignIn.JwtToken, common.ResourceIdParam{}, request.ResourceAttribute{})
	if err != nil {
		return
	}

	log.Info(".get resource attributes")
	resourceAttrs, _ := attributes.GetResourceAttrs(userSignIn.JwtToken)
	if err != nil {
		return
	}

	fmt.Println(resourceAttrs)
}
