package main

import (
	"development/go/recipes/lib/go-status/helpers/recipes"

	go_status "development/go/recipes/lib/go-status"

	log "github.com/sirupsen/logrus"
)

func RegisterFewUsers() {
	log.Info("Objective: Sign up a few new users")
	_, err := recipes.RegisterUser(go_status.DefaultEmail)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(go_status.ActiveUser)
	if err != nil {
		return
	}
}
