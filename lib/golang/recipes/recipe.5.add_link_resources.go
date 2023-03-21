package main

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/endpoints/block_pods.1"
	blocks "development/go/recipes/lib/golang/endpoints/blocks.1"
	keyPods "development/go/recipes/lib/golang/endpoints/key_pods.1"
	keys "development/go/recipes/lib/golang/endpoints/keys.1"
	registration "development/go/recipes/lib/golang/endpoints/registration.1"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	log "github.com/sirupsen/logrus"
)

const (
	CustomKeyType    = "CustomKey"
	Key1Name         = "CKey1"
	AnotherKeyName   = "CKey2"
	Block1Name       = "CBlock1"
	AnotherBlockName = "CBlock2"
	Pod1Name         = "CKeyPod1"
	BlockPod1Name    = "CBlockPod1"
)

// Add block, pod & block pod to a key and link them into another key
func main() {
	const Email = "api_code_user1@yopmail.com"
	user, err := signInToAddAndLink(Email, golang.Password)
	if err != nil {
		return
	}

	var newKey response.Key
	newKey, err = addCustomKey(user, Key1Name)
	if err != nil {
		return
	}

	var newPod response.Pod
	var newBlock response.Block
	var newBlockPod response.Pod
	newPod, newBlock, newBlockPod, err = addResources(user, newKey)
	if err != nil {
		return
	}

	log.Info("Add another key to link block & pod")
	helpers.SleepBefore()
	var anotherKey response.Key
	anotherKey, err = addCustomKey(user, AnotherKeyName)
	if err != nil {
		return
	}

	log.Info("Add another block to link pod")
	helpers.SleepBefore()
	var anotherBlock response.Block
	anotherBlock, err = addBlock(user, AnotherBlockName, newKey)
	if err != nil {
		return
	}

	err = linkResources(user, anotherKey, anotherBlock, newBlock, newBlockPod, newPod)
	if err != nil {
		return
	}
}

func linkResources(
	user response.User,
	anotherKey response.Key,
	anotherBlock response.Block,
	newBlock response.Block,
	newBlockPod response.Pod,
	newPod response.Pod,
) error {
	log.Info("Link key pod into this another key")
	helpers.SleepBefore()
	err := keyPods.LinkPodToKey(user.JwtToken, common.ResourceIdParam{
		PodId: newBlockPod.ID,
		KeyId: anotherKey.ID,
	})
	if err != nil {
		return err
	}
	log.Printf(".Block Pod, %s is linked successfully to %s Key.", newBlockPod.Name, anotherKey.Name)
	helpers.SleepAfter()

	log.Info("Link block into this another key")
	helpers.SleepBefore()
	err = blocks.LinkBlockToKey(user.JwtToken, common.ResourceIdParam{
		BlockId: newBlock.ID,
		KeyId:   anotherKey.ID,
	})
	if err != nil {
		return err
	}
	log.Printf(".Block, %s is linked successfully to %s Key.", newBlock.Name, anotherKey.Name)
	helpers.SleepAfter()

	log.Info("Link key pod into this another block")
	helpers.SleepBefore()
	err = block_pods.LinkPodToBlock(user.JwtToken, common.ResourceIdParam{
		PodId:   newPod.ID,
		BlockId: anotherBlock.ID,
		KeyId:   anotherKey.ID,
	})
	if err != nil {
		return err
	}
	log.Printf(".Pod, %s is linked successfully to %s Block.", newPod.Name, anotherBlock.Name)
	helpers.SleepAfter()
	return nil
}

func addResources(user response.User, newKey response.Key) (response.Pod, response.Block, response.Pod, error) {
	log.Info("Add a new key pod into this key")
	helpers.SleepBefore()
	newPod, err := keyPods.AddKeyPod(user.JwtToken, request.AddPodReqBody{
		Name: Pod1Name,
	}, newKey.ID)
	if err != nil {
		return response.Pod{}, response.Block{}, response.Pod{}, err
	}
	log.Printf(".Key Pod, %s is created successfully in %s Key.", newPod.Name, newKey.Name)
	helpers.SleepAfter()

	var newBlock response.Block
	newBlock, err = addBlock(user, Block1Name, newKey)
	if err != nil {
		return response.Pod{}, response.Block{}, response.Pod{}, err
	}

	log.Info("Add a new block pod into this block")
	helpers.SleepBefore()
	var newBlockPod response.Pod
	newBlockPod, err = block_pods.AddBlockPod(user.JwtToken, request.AddPodReqBody{
		Name: BlockPod1Name,
	}, common.ResourceIdParam{
		BlockId: newBlock.ID,
		KeyId:   newKey.ID,
	})
	if err != nil {
		return response.Pod{}, response.Block{}, response.Pod{}, err
	}
	log.Printf(".Block Pod, %s is created successfully in %s Block.", newBlockPod.Name, newBlock.Name)
	helpers.SleepAfter()
	return newPod, newBlock, newBlockPod, nil
}

func addBlock(user response.User, blockName string, newKey response.Key) (response.Block, error) {
	log.Info("Add a new block into this key")
	helpers.SleepBefore()
	newBlock, err := blocks.AddBlock(user.JwtToken, request.AddBlockReqBody{
		Name: blockName,
	}, newKey.ID)
	if err != nil {
		return response.Block{}, err
	}
	log.Printf(".Block, %s is created successfully in %s Key.", newBlock.Name, newKey.Name)
	helpers.SleepAfter()
	return newBlock, nil
}

func addCustomKey(user response.User, keyName string) (response.Key, error) {
	log.Info("Add a new custom key")
	helpers.SleepBefore()
	newKey, err := keys.AddKey(user.JwtToken, request.AddKeyReqBody{
		Name: keyName,
		Type: CustomKeyType,
	})
	if err != nil {
		return response.Key{}, err
	}
	log.Printf(".Key, %s is created successfully.", newKey.Name)
	helpers.SleepAfter()
	return newKey, nil
}

func signInToAddAndLink(email string, password string) (response.User, error) {
	log.Info("Sign in user with email: ", email)
	helpers.SleepBefore()
	user, err := registration.SignInByEmail(request.SignInReqBody{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return response.User{}, err
	}
	log.Info(".User successfully signed in, acquired JWT token")
	helpers.SleepAfter()
	return user, nil
}
