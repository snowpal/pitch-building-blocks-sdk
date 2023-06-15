package recipes

import (
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
	log.Info("- %s | %s", email, user.JwtToken)
}

func displayAllKeys(user response.User) {
	keys, err := keys.GetKeys(user.JwtToken, 0)
	if err != nil {
		return
	}
	for kIndex, key := range keys {
		log.Info("- %d. %s | %s", kIndex+1, key.Type, key.Name)

		blocks, err := blocks.GetBlocks(user.JwtToken, request.GetBlocksParam{
			KeyId: key.ID,
		})
		if err != nil {
			return
		}

		for bIndex, block := range blocks {
			log.Info("    - %d. %s", bIndex+1, block.Name)

			blockPods, err := blockPods.GetBlockPods(user.JwtToken, request.GetPodsParam{
				KeyId:   key.ID,
				BlockId: &block.ID,
			})
			if err != nil {
				return
			}

			for bpIndex, blockPod := range blockPods {
				log.Info("        - %d. %s", bpIndex+1, blockPod.Name)
			}
		}

		pods, err := keyPods.GetKeyPods(user.JwtToken, request.GetPodsParam{
			KeyId: key.ID,
		})
		if err != nil {
			return
		}

		for pIndex, pod := range pods {
			log.Info("    - %d. %s", pIndex+1, pod.Name)
		}
	}
}

func displayAllNotifications(user response.User) {
	notifications, err := notifications.GetNotifications(user.JwtToken)
	if err != nil {
		return
	}

	for index, notification := range notifications {
		log.Info("- %d. %s", index+1, notification.Text)
	}
}

func DisplayContent(user response.User, anotherUserEmail string) {
	log.Info("## Registered Users")
	displayUser(user.Email)
	displayUser(anotherUserEmail)

	log.Info("## Resources Created for user: %s", user.Email)
	displayAllKeys(user)

	anotherUser, err := recipes.SignIn(anotherUserEmail, lib.Password)
	if err != nil {
		return
	}

	log.Info("## Notifications for shared content as user: %s", anotherUserEmail)
	displayAllNotifications(anotherUser)
}
