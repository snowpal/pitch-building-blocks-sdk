package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/blocks/blocks.1"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/collaboration/collaboration.1.blocks"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/notifications"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	user2 "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/user"
)

const (
	KeyName          = "Diwali Festival"
	BlockName        = "Diwali Function"
	UpdatedBlockName = "Diwali Celebration"
)

func ShareBlock() {
	log.Info("Objective: Create block, share users as read & write, make 1 of them as admin.")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
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

	writeUser, err = recipes.SignIn(writeUser.Email, lib.Password)
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
	err = recipes.SearchUserAndShareBlock(user, block, "api_read_user", lib.ReadAcl)
	if err != nil {
		return block, err
	}
	err = recipes.SearchUserAndShareBlock(user, block, "api_write_user", lib.WriteAcl)
	if err != nil {
		return block, err
	}
	return block, nil
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
	systemKeys, _ := keys.GetKeysFilteredByType(writeUser.JwtToken, SystemKeyType)
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
