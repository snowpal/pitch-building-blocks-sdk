package recipes

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/endpoints/block_pods.1"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"fmt"
	"time"

	blocks "development/go/recipes/lib/golang/endpoints/blocks.1"
	collaboration "development/go/recipes/lib/golang/endpoints/collaboration.1.blocks"
	keyPods "development/go/recipes/lib/golang/endpoints/key_pods.1"
	keys "development/go/recipes/lib/golang/endpoints/keys.1"
)

func sleepWindow(sleepTime time.Duration) {
	time.Sleep(time.Second * sleepTime)
}

func SleepBefore() {
	sleepWindow(1)
}

func SleepAfter() {
	sleepWindow(2)
}

// ValidateDependencies We require that the first recipe be run before anything else as it registers a bunch of users.
// To verify if it was actually run, we do this "random" check.
func ValidateDependencies() (response.User, error) {
	user, err := SignIn(golang.DefaultEmail, golang.Password)
	fmt.Println(user)
	fmt.Println(err)
	if err != nil {
		return user, err
	}

	return user, nil
}

func addKey(user response.User, keyName string, keyType string) (response.Key, error) {
	newKey, err := keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: keyName,
			Type: keyType,
		})
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddCustomKey(user response.User, keyName string) (response.Key, error) {
	newKey, err := addKey(user, keyName, golang.CustomKeyType)
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddTeacherKey(user response.User, keyName string) (response.Key, error) {
	newKey, err := addKey(user, keyName, golang.TeacherKeyType)
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddProjectKey(user response.User, keyName string) (response.Key, error) {
	newKey, err := addKey(user, keyName, golang.ProjectKeyType)
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddBlock(user response.User, blockName string, key response.Key) (response.Block, error) {
	newBlock, err := blocks.AddBlock(
		user.JwtToken,
		request.AddBlockReqBody{Name: blockName},
		key.ID)
	if err != nil {
		return newBlock, err
	}
	return newBlock, nil
}

func AddPod(user response.User, podName string, key response.Key) (response.Pod, error) {
	newPod, err := keyPods.AddKeyPod(
		user.JwtToken,
		request.AddPodReqBody{Name: podName},
		key.ID)
	if err != nil {
		return newPod, err
	}
	return newPod, nil
}

func AddPodToBlock(user response.User, podName string, block response.Block) (response.Pod, error) {
	newPod, err := block_pods.AddBlockPod(
		user.JwtToken,
		request.AddPodReqBody{Name: podName},
		common.ResourceIdParam{BlockId: block.ID, KeyId: block.Key.ID})
	if err != nil {
		return newPod, err
	}
	return newPod, nil
}

func SearchUserAndShareBlock(user response.User, block response.Block, searchToken string, acl string) error {
	blockIdParam := common.ResourceIdParam{
		BlockId: block.ID,
		KeyId:   block.Key.ID,
	}

	searchedUsers, err := collaboration.GetUsersThisBlockCanBeSharedWith(
		user.JwtToken,
		common.SearchUsersParam{
			SearchToken: searchToken,
			ResourceIds: blockIdParam,
		})
	if err != nil {
		return err
	}

	// For the purpose of this recipe, it does not matter which user from the list we end up picking, hence we go with
	// the first one.
	_, err = collaboration.ShareBlockWithCollaborator(
		user.JwtToken,
		request.BlockAclReqBody{Acl: acl},
		common.AclParam{
			UserId:      searchedUsers[0].ID,
			ResourceIds: blockIdParam,
		})
	if err != nil {
		return err
	}
	return nil
}
