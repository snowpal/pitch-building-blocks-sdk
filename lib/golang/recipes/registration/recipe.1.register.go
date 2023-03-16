package main

import (
	"development/go/recipes/lib/golang"
	blocks "development/go/recipes/lib/golang/endpoints/blocks.1"
	keys "development/go/recipes/lib/golang/endpoints/keys.1"
	registration "development/go/recipes/lib/golang/endpoints/registration.1"
	"development/go/recipes/lib/golang/structs/request"
	log "github.com/sirupsen/logrus"

	"fmt"
)

// Sign up, activate user, sign in, get all keys.
func main() {
	log.Info(".sign up user with email: ", golang.Email)
	email := golang.Email
	password := golang.Password
	reqBody := request.UserRegistrationReqBody{
		Email:           &email,
		Password:        password,
		ConfirmPassword: &password,
	}
	user, err := registration.RegisterNewUserByEmail(reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Info(".activate user ID: ", user.ID)
	err = registration.ActivateUser(user.ID)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Info(".sign in user ID: ", user.ID)
	reqBody = request.UserRegistrationReqBody{
		Email:    &email,
		Password: password,
	}

	user, err = registration.SignInByEmail(reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Info(".get all keys: ", user.ID)
	allKeys, err := keys.GetKeys(user.JwtToken, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Info(allKeys)

	block, err := blocks.AddBlockBasedOnTemplate(user.JwtToken,
		request.AddBlockReqBody{Name: "Blk by Template1"},
		blocks.BlockByTemplateParam{KeyId: "63ceb29edb035900138d975d", TemplateId: "63c9bec50e98a800140720e3"})
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Info(block)
}
