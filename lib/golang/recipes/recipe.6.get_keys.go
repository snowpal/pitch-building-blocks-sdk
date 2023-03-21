package main

import (
	keys "development/go/recipes/lib/golang/endpoints/keys.1"
	"development/go/recipes/lib/golang/helpers/recipes"
	"development/go/recipes/lib/golang/structs/response"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Objective: Sign up, activate user, sign in, get all keys")
	user, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	//log.Info("Sign up user, email: ", golang.DefaultEmail)
	//var user response.User
	//user, err = recipes.RegisterUser(golang.DefaultEmail)
	//if err != nil {
	//	return
	//}

	log.Info("Get all keys: ", user.ID)
	var allKeys []response.Key
	allKeys, err = keys.GetKeys(user.JwtToken, 0)
	if err != nil {
		return
	}
	log.Info(allKeys)
}
