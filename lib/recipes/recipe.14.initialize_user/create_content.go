package recipes

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

var KeyNames = map[lib.KeyType]string{
	lib.Custom:  "Investments",
	lib.Teacher: "Math Class",
	lib.Project: "Snowpal Pitch",
}

var BlockNames = map[lib.KeyType]string{
	lib.Custom:  "Mutual Funds",
	lib.Teacher: "Algebra",
	lib.Project: "Mobile App",
}

var KeyPodName = "Fixed Diposite"

var BlockPodNames = map[lib.KeyType]string{
	lib.Custom:  "Systematic Investment Plan",
	lib.Teacher: "Sequence and Series",
	lib.Project: "Android App",
}

func createBlockWithPods(user response.User, keyType lib.KeyType, key response.Key) (BlockWithPods, error) {
	var blockWithPods BlockWithPods

	log.Info("Creating block inside ", key.Name, " key.")
	block, err := recipes.AddBlock(user, BlockNames[keyType], key)
	if err != nil {
		return blockWithPods, err
	}
	blockWithPods.Block = block

	log.Info("Creating block pod inside ", block.Name, " block.")
	var blockPod response.Pod
	blockPod, err = recipes.AddPodToBlock(user, BlockPodNames[lib.Custom], block)
	if err != nil {
		return blockWithPods, err
	}
	blockWithPods.BlockPods = append(blockWithPods.BlockPods, blockPod)

	return blockWithPods, err
}

func createKeyWithBlocks(user response.User, keyType lib.KeyType) (KeyWithResources, error) {
	var keyWithResources KeyWithResources

	log.Info("Creating custom key")
	key, err := recipes.AddKey(user, KeyNames[keyType], lib.KeyTypes[keyType])
	if err != nil {
		return keyWithResources, err
	}

	blockWithPods, err := createBlockWithPods(user, keyType, key)
	if err != nil {
		return keyWithResources, err
	}

	keyWithResources.Key = key
	keyWithResources.Blocks = append(keyWithResources.Blocks, blockWithPods)

	return keyWithResources, err
}

func createCustomKeyWithResources(user response.User) (KeyWithResources, error) {
	keyWithResources, err := createKeyWithBlocks(user, lib.Custom)
	if err != nil {
		return keyWithResources, err
	}

	log.Info("Creating key pod inside ", keyWithResources.Key.Name, " key.")
	var keyPod response.Pod
	keyPod, err = recipes.AddPod(user, KeyPodName, keyWithResources.Key)
	if err != nil {
		return keyWithResources, err
	}
	keyWithResources.Pods = append(keyWithResources.Pods, keyPod)

	return keyWithResources, err
}

func createTeacherKeyWithBlocks(user response.User) (KeyWithResources, error) {
	keyWithResources, err := createKeyWithBlocks(user, lib.Teacher)
	if err != nil {
		return keyWithResources, err
	}
	return keyWithResources, err
}

func createProjectKeyWithBlocks(user response.User) (KeyWithResources, error) {
	keyWithResources, err := createKeyWithBlocks(user, lib.Project)
	if err != nil {
		return keyWithResources, err
	}
	return keyWithResources, err
}

func CreateContent(user response.User) (AllKeys, error) {
	var allKeys AllKeys
	var err error

	log.Info("Creating custom key content")
	var customKey KeyWithResources
	customKey, err = createCustomKeyWithResources(user)
	if err != nil {
		return allKeys, err
	}
	allKeys.CustomKey = customKey

	log.Info("Creating teacher key content")
	var teacherKey KeyWithResources
	teacherKey, err = createTeacherKeyWithBlocks(user)
	if err != nil {
		return allKeys, err
	}
	allKeys.TeacherKey = teacherKey

	log.Info("Creating project key content")
	var projectKey KeyWithResources
	projectKey, err = createProjectKeyWithBlocks(user)
	if err != nil {
		return allKeys, err
	}
	allKeys.ProjectKey = projectKey

	return allKeys, err
}
