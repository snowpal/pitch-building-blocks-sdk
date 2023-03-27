package recipes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/endpoints/blocks/blocks.1"
	"development/go/recipes/lib/endpoints/collaboration/collaboration.1.blocks"
	"development/go/recipes/lib/endpoints/keys/keys.1"
	"development/go/recipes/lib/endpoints/notifications"
	"development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/request"
	"fmt"

	user2 "development/go/recipes/lib/endpoints/user"
	recipes2 "development/go/recipes/lib/helpers/recipes"
	response2 "development/go/recipes/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

const (
	KeyName          = "Diwali Festival"
	BlockName        = "Diwali Function"
	UpdatedBlockName = "Diwali Celebration"
)

func ShareBlock() {
	log.Info("Objective: Create block, share users as read & write, make 1 of them as admin.")
	_, err := recipes2.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes2.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	log.Info("Share a block")
	recipes2.SleepBefore()
	var block response2.Block
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
	recipes2.SleepBefore()
	err = showNotificationsAsWriteUser(writeUser)
	if err != nil {
		return
	}
	log.Printf(".Notifications for the recent share displayed successfully")
	recipes2.SleepAfter()

	log.Printf("Update block name as a write user")
	recipes2.SleepBefore()
	var resBlock response2.Block
	resBlock, err = updateBlockAsWriteUser(writeUser, block)
	if err != nil {
		return
	}
	log.Printf(".Write user updated block name to %s successfully", resBlock.Name)
	recipes2.SleepAfter()

	log.Printf("Grant admin access to a user with read access")
	err = makeReadUserAsAdmin(user, block)
	if err != nil {
		return
	}
	log.Printf(".Admin access has been granted successfully")
}

func getWriteUser(user response2.User, block response2.Block) (response2.User, error) {
	var writeUser response2.User
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

	writeUser, err = recipes2.SignIn(writeUser.Email, lib.Password)
	if err != nil {
		return writeUser, err
	}
	return writeUser, nil
}

func shareBlock(user response2.User) (response2.Block, error) {
	var block response2.Block
	key, err := recipes2.AddCustomKey(user, KeyName)
	if err != nil {
		return block, err
	}
	block, err = recipes2.AddBlock(user, BlockName, key)
	if err != nil {
		return block, err
	}
	err = recipes2.SearchUserAndShareBlock(user, block, "api_read_user", lib.ReadAcl)
	if err != nil {
		return block, err
	}
	err = recipes2.SearchUserAndShareBlock(user, block, "api_write_user", lib.WriteAcl)
	if err != nil {
		return block, err
	}
	return block, nil
}

func showNotificationsAsWriteUser(writeUser response2.User) error {
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

func updateBlockAsWriteUser(writeUser response2.User, block response2.Block) (response2.Block, error) {
	const (
		SystemKeyType       = "system"
		customSystemKeyType = "SharedCustomKey"
	)
	systemKeys, err := keys.GetKeysFilteredByType(writeUser.JwtToken, SystemKeyType)
	var customSystemKey response2.Key
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

func makeReadUserAsAdmin(user response2.User, block response2.Block) error {
	resBlock, err := collaboration.GetBlockCollaborators(
		user.JwtToken,
		common.ResourceIdParam{
			BlockId: block.ID,
			KeyId:   block.Key.ID,
		})
	if err != nil {
		return err
	}

	var readUser response2.SharedUser
	for _, sharedUser := range *resBlock.SharedUsers {
		if sharedUser.Acl == lib.ReadAcl {
			readUser = sharedUser
			break
		}
	}

	_, err = collaboration.UpdateBlockAcl(
		user.JwtToken,
		request.BlockAclReqBody{Acl: lib.AdminAcl},
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
