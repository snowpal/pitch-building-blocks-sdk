package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func ShareData(user response.User, anotherUserEmail string, allKeys AllKeys) error {
	log.Info(fmt.Sprintf("Share custom key data with %s", anotherUserEmail))
	err := shareCustomKeyData(user, anotherUserEmail, allKeys.CustomKey)
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Share teacher key data with %s", anotherUserEmail))
	err = shareClassroomKeyData(user, anotherUserEmail, allKeys.TeacherKey)
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Share project key data with %s", anotherUserEmail))
	err = shareProjectKeyData(user, anotherUserEmail, allKeys.ProjectKey)
	if err != nil {
		return err
	}
	return nil
}

func shareCustomKeyData(user response.User, anotherUserEmail string, customKey KeyWithResources) error {
	block := customKey.Blocks[0].Block
	log.Info(fmt.Sprintf("Share block %s for custom key, %s", block.Name, customKey.Key.Name))
	err := recipes.SearchUserAndShareBlock(user, block, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}

	blockPod := customKey.Blocks[0].BlockPods[0]
	log.Info(fmt.Sprintf("Share block pod %s for custom key, %s  and for %s block", blockPod.Name, customKey.Key.Name, block.Name))
	err = recipes.SearchUserAndShareBlockPod(user, blockPod, anotherUserEmail, lib.WriteAcl)
	if err != nil {
		return err
	}

	pod := customKey.Pods[0]
	log.Info(fmt.Sprintf("Share key pod %s for custom key, %s", pod.Name, customKey.Key.Name))
	err = recipes.SearchUserAndSharePod(user, pod, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}
	return nil
}

func shareClassroomKeyData(user response.User, anotherUserEmail string, teacherKey KeyWithResources) error {
	teacherBlock := teacherKey.Blocks[0].Block
	log.Info(fmt.Sprintf("Share block %s for teacher key, %s", teacherBlock.Name, teacherKey.Key.Name))
	err := recipes.SearchUserAndShareBlock(user, teacherBlock, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}
	return nil
}

func shareProjectKeyData(user response.User, anotherUserEmail string, projectKey KeyWithResources) error {
	projectBlock := projectKey.Blocks[0].Block
	log.Info(fmt.Sprintf("Share block %s for project key, %s", projectBlock.Name, projectKey.Key.Name))
	err := recipes.SearchUserAndShareBlock(user, projectBlock, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}
	return nil
}
