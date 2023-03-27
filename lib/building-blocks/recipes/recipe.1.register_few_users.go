package main

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers/recipes"
	log "github.com/sirupsen/logrus"
)

func RegisterFewUsers() {
	log.Info("Objective: Sign up a few new users")
	_, err := recipes.RegisterUser(building_blocks.DefaultEmail)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(building_blocks.ActiveUser)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(building_blocks.AdminUser)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(building_blocks.ReadUser)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(building_blocks.WriteUser)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(building_blocks.User1Email)
	if err != nil {
		return
	}
}
