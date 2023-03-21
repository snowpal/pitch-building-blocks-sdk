package main

import (
	"development/go/recipes/lib/golang"
	registration "development/go/recipes/lib/golang/endpoints/registration.1"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
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
	user, err := signInToShareBlock(ActiveUser, golang.Password)
	if err != nil {
		return
	}
	err = shareBlock(user)
	if err != nil {
		return
	}

	var writeUser response.User
	writeUser, err = signInToShareBlock(UserWithWrite, golang.Password)
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

func signInToShareBlock(email string, password string) (response.User, error) {
	log.Info("Sign in user with email: ", email)
	helpers.SleepBefore()
	user, err := registration.SignInByEmail(request.SignInReqBody{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return response.User{}, err
	}
	log.Info(".User successfully signed in, acquired JWT token")
	helpers.SleepAfter()
	return user, nil
}
