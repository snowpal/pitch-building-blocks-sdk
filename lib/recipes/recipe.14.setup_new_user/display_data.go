package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/blocks/blocks.1"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/notifications"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	blockPods "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/block_pods/block_pods.1"
	keyPods "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/key_pods/key_pods.1"
)

func displayUser(email string) {
	user, err := recipes.SignIn(email, lib.Password)
	if err != nil {
		return
	}
	log.Info(fmt.Sprintf("- %s | %s", email, user.JwtToken))
}

func displayAllKeys(user response.User) {
	allKeys, err := keys.GetKeys(user.JwtToken, 0)
	if err != nil {
		return
	}

	log.Info("List of Keys")
	for kIndex, key := range allKeys {
		if key.Type == lib.KeyTypes[lib.SharedCustom] ||
			key.Type == lib.KeyTypes[lib.SharedTeacher] ||
			key.Type == lib.KeyTypes[lib.SharedStudent] ||
			key.Type == lib.KeyTypes[lib.SharedProject] {
			continue
		}

		log.Info(fmt.Sprintf("%d. %s | %s", kIndex+1, key.Name, key.Type))

		var allBlocks []response.Block
		allBlocks, err = blocks.GetBlocks(user.JwtToken, request.GetBlocksParam{
			KeyId: key.ID,
		})
		if err != nil {
			return
		}

		log.Info(fmt.Sprintf("List of Blocks inside %s", key.Name))
		for bIndex, block := range allBlocks {
			log.Info(fmt.Sprintf("%d. %s", bIndex+1, block.Name))

			var allBlockPods []response.Pod
			allBlockPods, err = blockPods.GetBlockPods(user.JwtToken, request.GetPodsParam{
				KeyId:   key.ID,
				BlockId: &block.ID,
			})
			if err != nil {
				return
			}

			log.Info(fmt.Sprintf("List of Block Pods inside %s and %s", block.Name, key.Name))
			for bpIndex, blockPod := range allBlockPods {
				log.Info(fmt.Sprintf("%d. %s", bpIndex+1, blockPod.Name))
			}
		}

		if key.Type != lib.KeyTypes[lib.Custom] {
			continue
		}

		var allPods []response.Pod
		allPods, err = keyPods.GetKeyPods(user.JwtToken, request.GetPodsParam{
			KeyId: key.ID,
		})
		if err != nil {
			return
		}

		log.Info(fmt.Sprintf("List of Key Pods inside %s", key.Name))
		for pIndex, pod := range allPods {
			log.Info(fmt.Sprintf("%d. %s", pIndex+1, pod.Name))
		}
	}
}

func displayAllNotifications(user response.User) {
	allNotifications, err := notifications.GetNotifications(user.JwtToken)
	if err != nil {
		return
	}

	for index, notification := range allNotifications {
		log.Info(fmt.Sprintf("%d. %s", index+1, notification.Text))
	}
}

func DisplayData(user response.User, anotherUserEmail string) {
	log.Info("## Registered Users")
	displayUser(user.Email)
	displayUser(anotherUserEmail)

	log.Info(fmt.Sprintf("## Resources Created for user: %s", user.Email))
	displayAllKeys(user)

	anotherUser, err := recipes.SignIn(anotherUserEmail, lib.Password)
	if err != nil {
		return
	}

	log.Info(fmt.Sprintf("## Notifications for shared content as user: %s", anotherUserEmail))
	displayAllNotifications(anotherUser)
}
