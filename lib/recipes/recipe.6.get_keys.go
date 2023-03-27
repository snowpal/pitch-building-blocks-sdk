package recipes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/endpoints/keys/keys.1"
	"development/go/recipes/lib/structs/response"

	recipes2 "development/go/recipes/lib/helpers/recipes"

	log "github.com/sirupsen/logrus"
)

func GetAllKeys() {
	log.Info("Objective: Sign up, activate user, sign in, get all keys")
	_, err := recipes2.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes2.SignIn(lib.ActiveUser, lib.Password)
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
