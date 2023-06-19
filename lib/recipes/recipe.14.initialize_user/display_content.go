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
	log.Info("- ", email, " | ", user.JwtToken)
}

func displayAllKeys(user response.User) {
	keys, err := keys.GetKeys(user.JwtToken, 0)
	if err != nil {
		return
	}
	log.Info("List of Keys")
	for kIndex, key := range keys {
		if key.Type == lib.KeyTypes[lib.SharedCustom] ||
			key.Type == lib.KeyTypes[lib.SharedTeacher] ||
			key.Type == lib.KeyTypes[lib.SharedStudent] ||
			key.Type == lib.KeyTypes[lib.SharedProject] {
			continue
		}

		log.Info(kIndex+1, ". ", key.Name, " | ", key.Type)
		blocks, err := blocks.GetBlocks(user.JwtToken, request.GetBlocksParam{
			KeyId: key.ID,
		})
		if err != nil {
			return
		}

		log.Info("List of Blocks inside ", key.Name)
		for bIndex, block := range blocks {
			log.Info(bIndex+1, ". ", block.Name)

			blockPods, err := blockPods.GetBlockPods(user.JwtToken, request.GetPodsParam{
				KeyId:   key.ID,
				BlockId: &block.ID,
			})
			if err != nil {
				return
			}

			log.Info("List of Block Pods inside ", block.Name, " and ", key.Name)
			for bpIndex, blockPod := range blockPods {
				log.Info(bpIndex+1, ". ", blockPod.Name)
			}
		}

		if key.Type == lib.KeyTypes[lib.Custom] {
			continue
		}

		pods, err := keyPods.GetKeyPods(user.JwtToken, request.GetPodsParam{
			KeyId: key.ID,
		})
		if err != nil {
			return
		}

		log.Info("List of Key Pods inside ", key.Name)
		for pIndex, pod := range pods {
			log.Info(pIndex+1, ". ", pod.Name)
		}
	}
}

func displayAllNotifications(user response.User) {
	notifications, err := notifications.GetNotifications(user.JwtToken)
	if err != nil {
		return
	}

	for index, notification := range notifications {
		log.Info(index+1, ". ", notification.Text)
	}
}

func DisplayContent(user response.User, anotherUserEmail string) {
	log.Info("## Registered Users")
	displayUser(user.Email)
	displayUser(anotherUserEmail)

	log.Info("## Resources Created for user: ", user.Email)
	displayAllKeys(user)

	anotherUser, err := recipes.SignIn(anotherUserEmail, lib.Password)
	if err != nil {
		return
	}

	log.Info("## Notifications for shared content as user: ", anotherUserEmail)
	displayAllNotifications(anotherUser)
}
