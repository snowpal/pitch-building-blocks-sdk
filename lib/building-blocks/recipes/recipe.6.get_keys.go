package main

import (
	"development/go/recipes/lib/building-blocks"
	keys "development/go/recipes/lib/building-blocks/endpoints/keys.1"
	"development/go/recipes/lib/building-blocks/helpers/recipes"
	"development/go/recipes/lib/building-blocks/structs/response"
	log "github.com/sirupsen/logrus"
)

func GetAllKeys() {
	log.Info("Objective: Sign up, activate user, sign in, get all keys")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(building_blocks.ActiveUser, building_blocks.Password)
	if err != nil {
		return
	}

	log.Info("Get all keys: ", user.ID)
	var allKeys []response.Key
	allKeys, err = keys.GetKeys(user.JwtToken, 0)
	if err != nil {
		return
	}
	log.Info(allKeys)
}
