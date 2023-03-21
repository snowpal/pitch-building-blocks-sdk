package main

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers/recipes"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Objective: Sign up a few new users")
	_, err := recipes.RegisterUser(golang.DefaultEmail)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(golang.ActiveUser)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(golang.AdminUser)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(golang.UserWithRead)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(golang.UserWithWrite)
	if err != nil {
		return
	}
}