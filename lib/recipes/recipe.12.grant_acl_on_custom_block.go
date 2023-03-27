package recipes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/endpoints/blocks/blocks.1"
	"development/go/recipes/lib/structs/request"

	recipes2 "development/go/recipes/lib/helpers/recipes"
	response2 "development/go/recipes/lib/structs/response"
	log "github.com/sirupsen/logrus"
)

const (
	CopyKeyName   = "Insurance"
	CopyBlockName = "Car Insurance"
)

func GrantAclOnCustomBlock() {
	log.Info("Objective: Add Custom Block, Share Block, Grant Read Access, Copy Block, Grant Admin Access")
	_, err := recipes2.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes2.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	key, err := recipes2.AddCustomKey(user, CopyKeyName)
	if err != nil {
		return
	}

	log.Info("Add custom block")
	recipes2.SleepBefore()
	block, err := recipes2.AddBlock(user, CopyBlockName, key)
	if err != nil {
		return
	}
	log.Printf(".Block %s added successfully", block.Name)
	recipes2.SleepAfter()

	log.Info("Share block with read access")
	recipes2.SleepBefore()
	err = recipes2.SearchUserAndShareBlock(user, block, "api_read_user", lib.ReadAcl)
	if err != nil {
		return
	}
	log.Printf(".Block %s shared with %s with read access level", block.Name, lib.ReadUser)
	recipes2.SleepAfter()

	log.Info("Copy block and see acl is not copied")
	recipes2.SleepBefore()
	anotherBlock, err := copyBlock(user, block)
	if err != nil {
		return
	}
	log.Printf(".Block %s copied but %s don't have access on copied block", block.Name, lib.ReadUser)
	recipes2.SleepAfter()

	log.Info("Share block with admin access")
	recipes2.SleepBefore()
	err = recipes2.SearchUserAndShareBlock(user, anotherBlock, "api_admin_user", lib.AdminAcl)
	if err != nil {
		return
	}
	log.Printf(".Block %s shared with %s with admin access", block.Name, lib.ReadUser)
	recipes2.SleepAfter()
}

func copyBlock(user response2.User, block response2.Block) (response2.Block, error) {
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
