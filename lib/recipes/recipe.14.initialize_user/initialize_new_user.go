package recipes

import (
	log "github.com/sirupsen/logrus"
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func InitializeNewUser() {
	log.Info("Objective: Initialize new user with a dynamic email address and create content for that user")
	userEmail, err := RegisterNewUser()
	if err != nil {
		return
	}

	var user response.User
	user, err = recipes.SignIn(userEmail, lib.Password)
	if err != nil {
		return
	}

	var anotherUserEmail string
	anotherUserEmail, err = RegisterNewUser()
	if err != nil {
		return
	}

	var allKeys AllKeys
	allKeys, err = CreateContent(user)
	if err != nil {
		return
	}

	err = ShareContent(user, anotherUserEmail, allKeys)
	if err != nil {
		return
	}

	DisplayContent(user, anotherUserEmail)
}
