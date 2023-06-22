package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func SetupNewUser() {
	log.Info("Objective: Initialize new user with a dynamic email address and create data for that user")
	userEmail, err := RegisterNewUser()
	if err != nil {
		return
	}

	var user response.User
	user, err = recipes.SignIn(userEmail, lib.Password)
	if err != nil {
		return
	}

	log.Info(fmt.Sprintf("Creating data for %s", user.Email))
	var allKeys AllKeys
	allKeys, err = CreateData(user)
	if err != nil {
		return
	}

	log.Info("Register another user to share data")
	var anotherUserEmail string
	anotherUserEmail, err = RegisterNewUser()
	if err != nil {
		return
	}

	log.Info(fmt.Sprintf("Share data with %s", anotherUserEmail))
	err = ShareData(user, anotherUserEmail, allKeys)
	if err != nil {
		return
	}

	DisplayData(user, anotherUserEmail)
}
