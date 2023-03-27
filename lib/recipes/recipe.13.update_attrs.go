package recipes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/endpoints/attributes"
	recipes2 "development/go/recipes/lib/helpers/recipes"
	"development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/request"
	log "github.com/sirupsen/logrus"
)

const (
	AttrsKeyName   = "Birds"
	AttrsBlockName = "Parrot"
)

// UpdateAttributes sign in, update key attributes, update block attributes, update pod attributes, update block pod attributes,
// get resource attributes
func UpdateAttributes() {
	log.Info("Objective: Update show/hide of key, block, pod & block pod attributes")
	_, err := recipes2.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes2.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	log.Info("Get displayable attributes")
	recipes2.SleepBefore()
	resourceAttrs, _ := attributes.GetResourceAttrs(user.JwtToken)
	if err != nil {
		return
	}
	log.Info(resourceAttrs)

	log.Info("Update key attributes")
	recipes2.SleepBefore()
	key, err := recipes2.AddCustomKey(user, AttrsKeyName)
	if err != nil {
		return
	}
	err = attributes.UpdateKeyAttrs(
		user.JwtToken,
		key.ID,
		request.ResourceAttributeReqBody{
			AttributeNames: "tags,rendering_mode",
			ShowAttribute:  false,
		},
	)
	if err != nil {
		return
	}
	log.Printf(".Attributes for Key %s updated successfully", key.Name)
	recipes2.SleepAfter()

	log.Info("Update block attributes")
	recipes2.SleepBefore()
	block, err := recipes2.AddBlock(user, AttrsBlockName, key)
	if err != nil {
		return
	}
	err = attributes.UpdateBlockAttrs(
		user.JwtToken,
		common.ResourceIdParam{
			BlockId: block.ID,
			KeyId:   block.Key.ID,
		},
		request.ResourceAttributeReqBody{
			AttributeNames: "tags,rendering_mode",
			ShowAttribute:  false,
		},
	)
	if err != nil {
		return
	}
	log.Printf(".Attributes for block %s updated successfully", key.Name)
	recipes2.SleepAfter()
}
