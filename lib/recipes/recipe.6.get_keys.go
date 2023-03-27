package recipes

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	keys "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/keys/keys.1"
)

func GetAllKeys() {
	log.Info("Objective: Sign up, activate user, sign in, get all keys")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
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
