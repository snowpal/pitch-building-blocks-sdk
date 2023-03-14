package main

import (
	"development/go/recipes/lib/golang"
	blocks "development/go/recipes/lib/golang/endpoints/blocks.1"
	keys "development/go/recipes/lib/golang/endpoints/keys.1"
	registration "development/go/recipes/lib/golang/endpoints/registration.1"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	log "github.com/sirupsen/logrus"

	"fmt"
)

// Sign up, activate user, sign in, get all keys.
func main() {
	log.Info(".sign up user with email: ", golang.Email)
	userSignUp, err := registration.Signup(golang.Email)
	if err != nil {
		fmt.Println(err)
	}

	log.Info(".activate user ID: ", userSignUp.User.ID)
	registration.Activate(userSignUp.User.ID)

	log.Info(".sign in user ID: ", userSignUp.User.ID)
	var userSignIn response.UserRegistration
	userSignIn, err = registration.SignIn(golang.Email)
	if err != nil {
		fmt.Println(err)
	}

	log.Info(".get all keys: ", userSignIn.User.ID)
	allKeys, err := keys.GetKeys(userSignIn.User.JwtToken, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Info(allKeys)

	block, err := blocks.AddBlockBasedOnTemplate(userSignIn.User.JwtToken,
		request.AddBlockReqBody{Name: "Blk by Template1"},
		blocks.BlockByTemplateParam{KeyId: "63ceb29edb035900138d975d", TemplateId: "63c9bec50e98a800140720e3"})
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Info(block)
}
