package main

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers/recipes"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"

	blocks "development/go/recipes/lib/golang/endpoints/blocks.1"
	log "github.com/sirupsen/logrus"
)

const (
	CopyKeyName   = "Insurance"
	CopyBlockName = "Car Insurance"
)

func main() {
	log.Info("Objective: Add Custom Block, Share Block, Grant Read Access, Copy Block, Grant Admin Access")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(golang.ActiveUser, golang.Password)
	if err != nil {
		return
	}

	key, err := recipes.AddCustomKey(user, CopyKeyName)
	if err != nil {
		return
	}

	log.Info("Add custom block")
	recipes.SleepBefore()
	block, err := recipes.AddBlock(user, CopyBlockName, key)
	if err != nil {
		return
	}
	log.Printf(".Block %s added successfully", block.Name)
	recipes.SleepAfter()

	log.Info("Share block with read access")
	recipes.SleepBefore()
	err = recipes.SearchUserAndShareBlock(user, block, golang.ReadUserToken, golang.ReadAcl)
	if err != nil {
		return
	}
	log.Printf(".Block %s shared with %s with read access level", block.Name, golang.ReadUser)
	recipes.SleepAfter()

	log.Info("Copy block and see acl is not copied")
	recipes.SleepBefore()
	anotherBlock, err := copyBlock(user, block)
	if err != nil {
		return
	}
	log.Printf(".Block %s copied but %s don't have access on copied block", block.Name, golang.ReadUser)
	recipes.SleepAfter()

	log.Info("Share block with admin access")
	recipes.SleepBefore()
	err = recipes.SearchUserAndShareBlock(user, anotherBlock, golang.ReadUserToken, golang.AdminAcl)
	if err != nil {
		return
	}
	log.Printf(".Block %s shared with %s with admin access", block.Name, golang.ReadUser)
	recipes.SleepAfter()
}

func copyBlock(user response.User, block response.Block) (response.Block, error) {
	resBlock, err := blocks.CopyBlock(
		user.JwtToken,
		request.CopyMoveBlockParam{
			BlockId:       block.ID,
			KeyId:         block.Key.ID,
			TargetKeyId:   block.Key.ID,
			AllPods:       true,
			AllTasks:      true,
			AllChecklists: true,
		},
	)
	if err != nil {
		return resBlock, err
	}
	return resBlock, nil
}
