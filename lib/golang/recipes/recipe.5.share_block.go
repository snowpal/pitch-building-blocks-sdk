package main

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	log "github.com/sirupsen/logrus"
)

const (
	ActiveUser    = "apiuser3@gmail.com"
	UserWithRead  = "apiuser2@gmail.com"
	UserWithWrite = "api_code_user1@yopmail.com"
	AdminUser     = "apiuser1@gmail.com"
)

const (
	ShareTrialBlockName = "ShareTrailBlock"
)

func main() {
	log.Info("Objective: Create block, share users as read & write, make 1 of them as admin.")
	user, err := helpers.SignIn(ActiveUser, golang.Password)
	if err != nil {
		return
	}
	err = shareBlock(user)
	if err != nil {
		return
	}

	var writeUser response.User
	writeUser, err = helpers.SignIn(UserWithWrite, golang.Password)
	if err != nil {
		return
	}
	err = showNotifications(writeUser)
	if err != nil {
		return
	}

	err = updateBlockName(writeUser)
	if err != nil {
		return
	}

	err = makeReadUserAsAdmin(user)
	if err != nil {
		return
	}
}

func shareBlock(user response.User) error {
	return nil
}

func showNotifications(user response.User) error {
	return nil
}

func updateBlockName(user response.User) error {
	return nil
}

func makeReadUserAsAdmin(user response.User) error {
	return nil
}
