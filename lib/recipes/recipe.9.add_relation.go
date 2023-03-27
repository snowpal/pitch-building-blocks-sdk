package recipes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/endpoints/relations"
	recipes2 "development/go/recipes/lib/helpers/recipes"
	"development/go/recipes/lib/structs/request"
	response2 "development/go/recipes/lib/structs/response"
	log "github.com/sirupsen/logrus"
)

const (
	RelationKeyName   = "Animals"
	RelationBlockName = "Tiger"
	RelationPodName   = "Cat"
)

func AddRelation() {
	log.Info("Objective: Add and Remove Relations")
	_, err := recipes2.ValidateDependencies()
	if err != nil {
		return
	}

	log.Info("Create a key and block & pod into that key")
	user, err := recipes2.SignIn(lib.ActiveUser, lib.Password)
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

func removeRelation(user response2.User, block response2.Block, pod response2.Pod) error {
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

func addRelation(user response2.User) (response2.Block, response2.Pod, error) {
	var (
		block response2.Block
		pod   response2.Pod
	)
	key, err := recipes2.AddCustomKey(user, RelationKeyName)
	if err != nil {
		return block, pod, err
	}
	block, err = recipes2.AddBlock(user, RelationBlockName, key)
	if err != nil {
		return block, pod, err
	}
	pod, err = recipes2.AddPod(user, RelationPodName, key)
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
