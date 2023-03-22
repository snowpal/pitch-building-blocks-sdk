package main

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers/recipes"
	"development/go/recipes/lib/golang/structs/response"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Objective: Add Custom Block, Share Block, Grant Read Access, Copy Block, Grant Admin Access")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(golang.DefaultEmail, golang.Password)
	if err != nil {
		return
	}

	key, err := addKey(user, "")
	block, err := addBlock(user, "", key)
	grantAccess(user, block, "")
	anotherBlock, err := copyBlock(user, block, key)
	grantAccess(user, anotherBlock, "")
}

func copyBlock(user response.User, block response.Block, key response.Key) (response.Block, error) {
	return response.Block{}, nil
}

func grantAccess(user response.User, block response.Block, accessLevel string) {
}
