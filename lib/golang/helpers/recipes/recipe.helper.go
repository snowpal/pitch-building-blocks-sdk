package recipes

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"fmt"
	"time"

	blocks "development/go/recipes/lib/golang/endpoints/blocks.1"
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

func AddCustomKey(user response.User, keyName string) (response.Key, error) {
	newKey, err := keys.AddKey(
		user.JwtToken,
		request.AddKeyReqBody{
			Name: keyName,
			Type: golang.CustomKeyType,
		})
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
