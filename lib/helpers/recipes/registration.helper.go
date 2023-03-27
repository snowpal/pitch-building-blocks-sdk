package recipes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/endpoints/registration"
	"development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func RegisterUser(email string) (response.User, error) {
	signUpReqBody := request.SignupReqBody{
		Email:           email,
		Password:        lib.Password,
		ConfirmPassword: lib.Password,
	}
	user, err := registration.RegisterNewUserByEmail(signUpReqBody)
	if err != nil {
		fmt.Println(err)
		return response.User{}, err
	}

	log.Info(".activate user ID: ", user.ID)
	err = registration.ActivateUser(user.ID)
	if err != nil {
		fmt.Println(err)
		return response.User{}, err
	}
	return user, nil
}

func SignIn(email string, password string) (response.User, error) {
	log.Info("Sign in user with email: ", email)
	SleepBefore()
	user, err := registration.SignInByEmail(request.SignInReqBody{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return response.User{}, err
	}
	log.Info(".User successfully signed in, acquired JWT token")
	SleepAfter()
	return user, nil
}
