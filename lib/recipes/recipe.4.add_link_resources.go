package recipes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/endpoints/block_pods/block_pods.1"
	"development/go/recipes/lib/endpoints/blocks/blocks.1"
	"development/go/recipes/lib/endpoints/key_pods/key_pods.1"
	"development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/request"

	recipes2 "development/go/recipes/lib/helpers/recipes"
	response2 "development/go/recipes/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

const (
	Key1Name         = "Taxes"
	AnotherKeyName   = "State Taxes"
	Block1Name       = "Form 1040"
	AnotherBlockName = "Form 1120S"
	Pod1Name         = "Income"
	BlockPod1Name    = "Expenses"
)

// AddAndLinkResources Add block, pod & block pod to a key and link them into another key
func AddAndLinkResources() {
	log.Info("Objective: Add keys and blocks, and link blocks")
	_, err := recipes2.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes2.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	var newKey response2.Key
	log.Info("Add a new custom key")
	recipes2.SleepBefore()
	newKey, err = recipes2.AddCustomKey(user, Key1Name)
	if err != nil {
		return
	}
	log.Printf(".Key, %s is added successfully.", newKey.Name)
	recipes2.SleepAfter()

	var (
		newPod      response2.Pod
		newBlock    response2.Block
		newBlockPod response2.Pod
	)
	newPod, newBlock, newBlockPod, err = addBlocksAndPods(user, newKey)
	if err != nil {
		return
	}

	log.Info("Add another key")
	recipes2.SleepBefore()
	var anotherKey response2.Key
	anotherKey, err = recipes2.AddCustomKey(user, AnotherKeyName)
	if err != nil {
		return
	}

	log.Info("Add block")
	recipes2.SleepBefore()
	var anotherBlock response2.Block
	anotherBlock, err = recipes2.AddBlock(user, AnotherBlockName, newKey)
	if err != nil {
		return
	}
	log.Printf(".Block, %s is created successfully.", newBlock.Name)
	recipes2.SleepAfter()

	err = linkResources(user, anotherKey, anotherBlock, newBlock, newBlockPod, newPod)
	if err != nil {
		return
	}
}

func linkResources(
	user response2.User,
	anotherKey response2.Key,
	anotherBlock response2.Block,
	newBlock response2.Block,
	newBlockPod response2.Pod,
	newPod response2.Pod,
) error {
	log.Info("Link key pod into the other key")
	recipes2.SleepBefore()
	err := keyPods.LinkPodToKey(user.JwtToken,
		common.ResourceIdParam{
			PodId: newPod.ID,
			KeyId: anotherKey.ID,
		})
	if err != nil {
		return err
	}
	log.Printf(".Block Pod, %s is linked successfully to %s Key.", newPod.Name, anotherKey.Name)
	recipes2.SleepAfter()

	log.Info("Link block into the other key")
	recipes2.SleepBefore()
	err = blocks.LinkBlockToKey(user.JwtToken,
		common.ResourceIdParam{
			BlockId: newBlock.ID,
			KeyId:   anotherKey.ID,
		})
	if err != nil {
		return err
	}
	log.Printf(".Block, %s is linked successfully to %s Key.", newBlock.Name, anotherKey.Name)
	recipes2.SleepAfter()

	log.Info("Link key pod into the other block")
	recipes2.SleepBefore()
	err = blockPods.LinkPodToBlock(user.JwtToken,
		common.ResourceIdParam{
			PodId:   newBlockPod.ID,
			BlockId: anotherBlock.ID,
			KeyId:   anotherKey.ID,
		})
	if err != nil {
		return err
	}
	log.Printf(".Pod, %s is linked successfully to %s Block.", newPod.Name, anotherBlock.Name)
	recipes2.SleepAfter()
	return nil
}

func addBlocksAndPods(user response2.User, newKey response2.Key) (response2.Pod, response2.Block, response2.Pod, error) {
	log.Info("Add a new key pod into this key")
	recipes2.SleepBefore()

	var (
		pod   response2.Pod
		block response2.Block
	)
	newPod, err := keyPods.AddKeyPod(user.JwtToken,
		request.AddPodReqBody{
			Name: Pod1Name,
		},
		newKey.ID)
	if err != nil {
		return pod, block, pod, err
	}
	log.Printf(".Key Pod, %s is created successfully in %s Key.", newPod.Name, newKey.Name)
	recipes2.SleepAfter()

	log.Info("Add a new block")
	recipes2.SleepBefore()
	var newBlock response2.Block
	newBlock, err = recipes2.AddBlock(user, Block1Name, newKey)
	if err != nil {
		return pod, block, pod, err
	}
	log.Printf(".Block, %s is created successfully.", newBlock.Name)
	recipes2.SleepAfter()

	log.Info("Add a new block pod in this block")
	recipes2.SleepBefore()
	var newBlockPod response2.Pod
	newBlockPod, err = blockPods.AddBlockPod(user.JwtToken,
		request.AddPodReqBody{
			Name: BlockPod1Name,
		},
		common.ResourceIdParam{
			BlockId: newBlock.ID,
			KeyId:   newKey.ID,
		})
	if err != nil {
		return pod, block, pod, err
	}
	log.Printf(".Block Pod, %s is created successfully in %s Block.", newBlockPod.Name, newBlock.Name)
	recipes2.SleepAfter()
	return newPod, newBlock, newBlockPod, nil
}
