package recipes

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/blocks/blocks.1"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/relations"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"

	log "github.com/sirupsen/logrus"
	keyPods "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/key_pods/key_pods.1"
	recipes "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	response "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

const (
	RelationKeyName   = "Animals"
	RelationBlockName = "Tiger"
	RelationPodName   = "Cat"
)

func AddRelation() {
	log.Info("Objective: Add and Remove Relations")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	log.Info("Create a key and block & pod into that key")
	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	log.Info("Relate the block with key pod")
	block, pod, err := addRelation(user)
	if err != nil {
		return
	}
	log.Printf(".Block %s is related with pod %s successfully", block.Name, pod.Name)

	log.Info("Unrelate the block from key pod")
	err = removeRelation(user, block, pod)
	if err != nil {
		return
	}
	log.Printf(".Block %s is unrelated from pod %s successfully", block.Name, pod.Name)
}

func removeRelation(user response.User, block response.Block, pod response.Pod) error {
	err := relations.UnrelateBlockFromKeyPod(
		user.JwtToken,
		request.BlockToPodRelationParam{
			BlockId:     block.ID,
			TargetPodId: pod.ID,
			TargetKeyId: pod.Key.ID,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func addRelation(user response.User) (response.Block, response.Pod, error) {
	var (
		block response.Block
		pod   response.Pod
	)
	key, err := keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: RelationKeyName,
			Type: lib.KeyTypes[lib.Custom],
		})
	if err != nil {
		return block, pod, err
	}
	block, err = blocks.AddBlock(
		user.JwtToken,
		request.AddBlockReqBody{Name: RelationBlockName},
		key.ID)
	if err != nil {
		return block, pod, err
	}
	pod, err = keyPods.AddKeyPod(
		user.JwtToken,
		request.AddPodReqBody{Name: RelationPodName},
		key.ID,
	)
	if err != nil {
		return block, pod, err
	}
	err = relations.RelateBlockToKeyPod(
		user.JwtToken,
		request.BlockToPodRelationParam{
			BlockId:     block.ID,
			TargetPodId: pod.ID,
			TargetKeyId: pod.Key.ID,
		},
	)
	if err != nil {
		return block, pod, err
	}
	return block, pod, nil

}
