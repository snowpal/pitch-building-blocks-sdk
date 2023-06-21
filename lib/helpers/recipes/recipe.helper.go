package recipes

import (
	"time"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/blocks/blocks.1"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/keys/keys.1"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	blockPods "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/block_pods/block_pods.1"
	keyPods "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/key_pods/key_pods.1"
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
	user, err := SignIn(lib.DefaultEmail, lib.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func AddKey(user response.User, keyName string, keyType string) (response.Key, error) {
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
	newKey, err := AddKey(user, keyName, lib.KeyTypes[lib.Custom])
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddTeacherKey(user response.User, keyName string) (response.Key, error) {
	newKey, err := AddKey(user, keyName, lib.KeyTypes[lib.Teacher])
	if err != nil {
		return newKey, err
	}
	return newKey, nil
}

func AddProjectKey(user response.User, keyName string) (response.Key, error) {
	newKey, err := AddKey(user, keyName, lib.KeyTypes[lib.Project])
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
	newPod, err := blockPods.AddBlockPod(
		user.JwtToken,
		request.AddPodReqBody{Name: podName},
		common.ResourceIdParam{BlockId: block.ID, KeyId: block.Key.ID})
	if err != nil {
		return newPod, err
	}
	return newPod, nil
}
