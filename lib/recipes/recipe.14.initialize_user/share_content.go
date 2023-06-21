package recipes

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func shareBlock(user response.User, anotherUserEmail string, block BlockWithPods) error {
	err := recipes.SearchUserAndShareBlock(user, block.Block, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}
	return nil
}

func shareCustomBlockAndPod(user response.User, anotherUserEmail string, customKey KeyWithResources) error {
	log.Info("Share block ", customKey.Blocks[0].Block.Name, " for custom key, ", customKey.Key.Name)
	err := shareBlock(user, anotherUserEmail, customKey.Blocks[0])
	if err != nil {
		return err
	}

	log.Info("Share block pod ", customKey.Blocks[0].BlockPods[0].Name, " for custom key, ", customKey.Key.Name, " and for ", customKey.Blocks[0].Block.Name)
	err = recipes.SearchUserAndShareBlockPod(user, customKey.Blocks[0].BlockPods[0], anotherUserEmail, lib.WriteAcl)
	if err != nil {
		return err
	}

	log.Info("Share key pod ", customKey.Pods[0].Name, " for custom key, ", customKey.Key.Name)
	err = recipes.SearchUserAndSharePod(user, customKey.Pods[0], anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}
	return nil
}

func shareClassroomContent(user response.User, anotherUserEmail string, teacherKey KeyWithResources) error {
	log.Info("Share block ", teacherKey.Blocks[0].Block.Name, " for teacher key, ", teacherKey.Key.Name)
	err := shareBlock(user, anotherUserEmail, teacherKey.Blocks[0])
	if err != nil {
		return err
	}
	return nil
}

func shareProjectManagementContent(user response.User, anotherUserEmail string, projectKey KeyWithResources) error {
	log.Info("Share block ", projectKey.Blocks[0].Block.Name, " for project key, ", projectKey.Key.Name)
	err := shareBlock(user, anotherUserEmail, projectKey.Blocks[0])
	if err != nil {
		return err
	}
	return nil
}

func ShareContent(user response.User, anotherUserEmail string, allKeys AllKeys) error {
	log.Info("Share custom key content with ", anotherUserEmail)
	err := shareCustomBlockAndPod(user, anotherUserEmail, allKeys.CustomKey)
	if err != nil {
		return err
	}

	log.Info("Share teacher key content with ", anotherUserEmail)
	err = shareClassroomContent(user, anotherUserEmail, allKeys.TeacherKey)
	if err != nil {
		return err
	}

	log.Info("Share project key content with ", anotherUserEmail)
	err = shareProjectManagementContent(user, anotherUserEmail, allKeys.ProjectKey)
	if err != nil {
		return err
	}
	return nil
}
