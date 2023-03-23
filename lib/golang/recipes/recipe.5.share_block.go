package main

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/endpoints/notifications.1"
	"development/go/recipes/lib/golang/helpers/recipes"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"fmt"

	blocks "development/go/recipes/lib/golang/endpoints/blocks.1"
	collaboration "development/go/recipes/lib/golang/endpoints/collaboration.1.blocks"
	keys "development/go/recipes/lib/golang/endpoints/keys.1"
	user2 "development/go/recipes/lib/golang/endpoints/user.1"
	log "github.com/sirupsen/logrus"
)

const (
	KeyName          = "Diwali Festival11"
	BlockName        = "Diwali Function11"
	UpdatedBlockName = "Diwali Celebration11"
)

func main() {
	log.Info("Objective: Create block, share users as read & write, make 1 of them as admin.")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(golang.ActiveUser, golang.Password)
	if err != nil {
		return
	}

	log.Info("Share a block")
	recipes.SleepBefore()
	var block response.Block
	block, err = shareBlock(user)
	if err != nil {
		return
	}

	writeUser, err := getWriteUser(user, block)
	fmt.Println(user.JwtToken)
	if err != nil {
		return
	}

	log.Info("Show notifications as write user")
	recipes.SleepBefore()
	err = showNotificationsAsWriteUser(writeUser)
	if err != nil {
		return
	}
	log.Printf(".Notifications for the recent share displayed successfully")
	recipes.SleepAfter()

	log.Printf("Update block name as a write user")
	recipes.SleepBefore()
	var resBlock response.Block
	resBlock, err = updateBlockAsWriteUser(writeUser, block)
	if err != nil {
		return
	}
	log.Printf(".Write user updated block name to %s successfully", resBlock.Name)
	recipes.SleepAfter()

	log.Printf("Grant admin access to a user with read access")
	err = makeReadUserAsAdmin(user, block)
	if err != nil {
		return
	}
	log.Printf(".Admin access has been granted successfully")
}

func getWriteUser(user response.User, block response.Block) (response.User, error) {
	var writeUser response.User
	resBlock, err := collaboration.GetBlockCollaborators(
		user.JwtToken,
		common.ResourceIdParam{
			BlockId: block.ID,
			KeyId:   block.Key.ID,
		})
	if err != nil {
		return writeUser, err
	}
	allUsers, err := user2.GetUsers(user.JwtToken)
	for _, sharedUser := range *resBlock.SharedUsers {
		for _, userInAll := range allUsers {
			if sharedUser.ID == userInAll.ID {
				writeUser = userInAll
				break
			}
		}
	}
	if err != nil {
		return writeUser, err
	}

	writeUser, err = recipes.SignIn(writeUser.Email, golang.Password)
	if err != nil {
		return writeUser, err
	}
	return writeUser, nil
}

func shareBlock(user response.User) (response.Block, error) {
	var block response.Block
	key, err := recipes.AddCustomKey(user, KeyName)
	if err != nil {
		return block, err
	}
	block, err = recipes.AddBlock(user, BlockName, key)
	if err != nil {
		return block, err
	}
	err = searchUserAndShareBlock(user, block, golang.ReadUserToken, golang.ReadAcl)
	if err != nil {
		return block, err
	}
	err = searchUserAndShareBlock(user, block, golang.WriteUserToken, golang.WriteAcl)
	if err != nil {
		return block, err
	}
	return block, nil
}

func searchUserAndShareBlock(user response.User, block response.Block, searchToken string, acl string) error {
	blockIdParam := common.ResourceIdParam{
		BlockId: block.ID,
		KeyId:   block.Key.ID,
	}

	searchedUsers, err := collaboration.GetUsersThisBlockCanBeSharedWith(
		user.JwtToken,
		common.SearchUsersParam{
			SearchToken: searchToken,
			ResourceIds: blockIdParam,
		})
	if err != nil {
		return err
	}

	_, err = collaboration.ShareBlockWithCollaborator(
		user.JwtToken,
		request.BlockAclReqBody{Acl: acl},
		common.AclParam{
			UserId:      searchedUsers[0].ID,
			ResourceIds: blockIdParam,
		})
	if err != nil {
		return err
	}
	return nil
}

func showNotificationsAsWriteUser(writeUser response.User) error {
	unreadNotifications, err := notifications.GetNotifications(writeUser.JwtToken)
	if err != nil {
		return err
	}
	fmt.Println(len(unreadNotifications))
	for index, notification := range unreadNotifications {
		if notification.Type == "acl" {
			log.Printf(".Notification %d: %s", index, notification.Text)
		}
	}
	return nil
}

func updateBlockAsWriteUser(writeUser response.User, block response.Block) (response.Block, error) {
	const (
		SystemKeyType       = "system"
		customSystemKeyType = "SharedCustomKey"
	)
	systemKeys, err := keys.GetKeysFilteredByType(writeUser.JwtToken, SystemKeyType)
	var customSystemKey response.Key
	for _, systemKey := range systemKeys {
		if systemKey.Type == customSystemKeyType {
			customSystemKey = systemKey
			break
		}
	}
	updatedBlockName := UpdatedBlockName
	resBlock, err := blocks.UpdateBlock(
		writeUser.JwtToken,
		blocks.UpdateBlockReqBody{Name: &updatedBlockName},
		common.ResourceIdParam{
			BlockId: block.ID,
			KeyId:   customSystemKey.ID,
		})
	if err != nil {
		return resBlock, err
	}
	return resBlock, nil
}

func makeReadUserAsAdmin(user response.User, block response.Block) error {
	resBlock, err := collaboration.GetBlockCollaborators(
		user.JwtToken,
		common.ResourceIdParam{
			BlockId: block.ID,
			KeyId:   block.Key.ID,
		})
	if err != nil {
		return err
	}

	var readUser response.SharedUser
	for _, sharedUser := range *resBlock.SharedUsers {
		if sharedUser.Acl == golang.ReadAcl {
			readUser = sharedUser
			break
		}
	}

	_, err = collaboration.UpdateBlockAcl(
		user.JwtToken,
		request.BlockAclReqBody{Acl: golang.AdminAcl},
		common.AclParam{
			UserId: readUser.ID,
			ResourceIds: common.ResourceIdParam{
				BlockId: block.ID,
				KeyId:   block.Key.ID,
			},
		})
	if err != nil {
		return err
	}
	return nil
}