package recipes

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/attributes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"

	log "github.com/sirupsen/logrus"
	recipes "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
)

const (
	AttrsKeyName   = "Birds"
	AttrsBlockName = "Parrot"
)

// UpdateAttributes sign in, update key attributes, update block attributes, update pod attributes, update block pod attributes,
// get resource attributes
func UpdateAttributes() {
	log.Info("Objective: Update show/hide of key, block, pod & block pod attributes")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	log.Info("Get displayable attributes")
	recipes.SleepBefore()
	resourceAttrs, _ := attributes.GetResourceAttrs(user.JwtToken)
	if err != nil {
		return
	}
	log.Info(resourceAttrs)

	log.Info("Update key attributes")
	recipes.SleepBefore()
	key, err := recipes.AddCustomKey(user, AttrsKeyName)
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
	recipes.SleepAfter()

	log.Info("Update block attributes")
	recipes.SleepBefore()
	block, err := recipes.AddBlock(user, AttrsBlockName, key)
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
	recipes.SleepAfter()
}
