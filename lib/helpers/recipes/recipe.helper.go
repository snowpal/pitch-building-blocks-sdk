package recipes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/endpoints/block_pods/block_pods.1"
	"development/go/recipes/lib/endpoints/blocks/blocks.1"
	"development/go/recipes/lib/endpoints/collaboration/collaboration.1.blocks"
	"development/go/recipes/lib/endpoints/key_pods/key_pods.1"
	"development/go/recipes/lib/endpoints/keys/keys.1"
	"development/go/recipes/lib/structs/common"
	"fmt"
	"time"

	request2 "development/go/recipes/lib/structs/request"
	response2 "development/go/recipes/lib/structs/response"
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
func ValidateDependencies() (response2.User, error) {
	user, err := SignIn(lib.DefaultEmail, lib.Password)
	fmt.Println(user)
	fmt.Println(err)
	if err != nil {
		return user, err
	}

	return user, nil
}

func addKey(user response2.User, keyName string, keyType string) (response2.Key, error) {
	newKey, err := keys.AddKey(
		user.JwtToken,
		request2.AddKeyReqBody{
			Name: keyName,
			Type: keyType,
		})
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddCustomKey(user response2.User, keyName string) (response2.Key, error) {
	newKey, err := addKey(user, keyName, lib.CustomKeyType)
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddTeacherKey(user response2.User, keyName string) (response2.Key, error) {
	newKey, err := addKey(user, keyName, lib.TeacherKeyType)
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddProjectKey(user response2.User, keyName string) (response2.Key, error) {
	newKey, err := addKey(user, keyName, lib.ProjectKeyType)
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddBlock(user response2.User, blockName string, key response2.Key) (response2.Block, error) {
	newBlock, err := blocks.AddBlock(
		user.JwtToken,
		request2.AddBlockReqBody{Name: blockName},
		key.ID)
	if err != nil {
		return newBlock, err
	}
	return newBlock, nil
}

func AddPod(user response2.User, podName string, key response2.Key) (response2.Pod, error) {
	newPod, err := keyPods.AddKeyPod(
		user.JwtToken,
		request2.AddPodReqBody{Name: podName},
		key.ID)
	if err != nil {
		return newPod, err
	}
	return newPod, nil
}

func AddPodToBlock(user response2.User, podName string, block response2.Block) (response2.Pod, error) {
	newPod, err := blockPods.AddBlockPod(
		user.JwtToken,
		request2.AddPodReqBody{Name: podName},
		common.ResourceIdParam{BlockId: block.ID, KeyId: block.Key.ID})
	if err != nil {
		return newPod, err
	}
	return newPod, nil
}

func SearchUserAndShareBlock(user response2.User, block response2.Block, searchToken string, acl string) error {
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
		request2.BlockAclReqBody{Acl: acl},
		common.AclParam{
			UserId:      searchedUsers[0].ID,
			ResourceIds: blockIdParam,
		})
	if err != nil {
		return err
	}
	return nil
}
