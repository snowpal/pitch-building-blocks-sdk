package main

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/endpoints/relations.1"
	"development/go/recipes/lib/building-blocks/helpers/recipes"
	"development/go/recipes/lib/building-blocks/structs/request"
	"development/go/recipes/lib/building-blocks/structs/response"
	log "github.com/sirupsen/logrus"
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
	user, err := recipes.SignIn(building_blocks.ActiveUser, building_blocks.Password)
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
	key, err := recipes.AddCustomKey(user, RelationKeyName)
	if err != nil {
		return block, pod, err
	}
	block, err = recipes.AddBlock(user, RelationBlockName, key)
	if err != nil {
		return block, pod, err
	}
	pod, err = recipes.AddPod(user, RelationPodName, key)
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
