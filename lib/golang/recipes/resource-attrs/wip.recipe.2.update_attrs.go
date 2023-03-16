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
	log.Info(".sign in user email: ", golang.Email)
	userSignIn, err := registration.SignInByEmail(golang.Email)
	if err != nil {
		fmt.Println(err)
	}

	log.Info(".TODO(): get a key, and use it below")
	attributes.UpdateKeyAttrs(userSignIn.User.JwtToken, "", request.ResourceAttribute{})

	log.Info(".TODO(): get a block, and use it below")
	attributes.UpdateBlockAttrs(userSignIn.User.JwtToken, common.ResourceIdParam{}, request.ResourceAttribute{})

	log.Info(".TODO(): get a pod, and use it below")
	attributes.UpdatePodAttrs(userSignIn.User.JwtToken, common.ResourceIdParam{}, request.ResourceAttribute{})

	log.Info(".TODO(): get a block pod, and use it below")
	attributes.UpdateBlockPodAttrs(userSignIn.User.JwtToken, common.ResourceIdParam{}, request.ResourceAttribute{})

	log.Info(".get resource attributes")
	resourceAttrs, _ := attributes.GetResourceAttrs(userSignIn.User.JwtToken)
	if err != nil {
		return
	}

	fmt.Println(resourceAttrs)
}
