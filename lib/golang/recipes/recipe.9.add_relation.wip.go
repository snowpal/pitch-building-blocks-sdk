package main

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/endpoints/relations.1"
	"development/go/recipes/lib/golang/helpers/recipes"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	log "github.com/sirupsen/logrus"
)

const (
	RelationKeyName   = "Animals6"
	RelationBlockName = "Tiger6"
	RelationPodName   = "Cat6"
)

func main() {
	log.Info("Objective: Add and Remove Relations")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(golang.ActiveUser, golang.Password)
	if err != nil {
		return
	}

	block, pod, err := addRelation(user)
	if err != nil {
		return
	}
	err = removeRelation(user, block, pod)
	if err != nil {
		return
	}
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
