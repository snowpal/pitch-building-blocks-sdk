package recipes

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func shareBlockWithBlockPods(user response.User, anotherUserEmail string, block BlockWithPods) error {
	err := recipes.SearchUserAndShareBlock(user, block.Block, anotherUserEmail, lib.WriteAcl)
	if err != nil {
		return err
	}

	err = recipes.SearchUserAndShareBlockPod(user, block.BlockPods[0], anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}
	return nil
}

func shareProjectManagementContent(user response.User, anotherUserEmail string, projectKey KeyWithResources) error {
	err := shareBlockWithBlockPods(user, anotherUserEmail, projectKey.Blocks[0])
	if err != nil {
		return err
	}
	return nil
}

func shareClassroomContent(user response.User, anotherUserEmail string, teacherKey KeyWithResources) error {
	err := shareBlockWithBlockPods(user, anotherUserEmail, teacherKey.Blocks[0])
	if err != nil {
		return err
	}
	return nil
}

func shareCustomBlockAndPod(user response.User, anotherUserEmail string, customKey KeyWithResources) error {
	err := shareBlockWithBlockPods(user, anotherUserEmail, customKey.Blocks[0])
	if err != nil {
		return err
	}

	err = recipes.SearchUserAndSharePod(user, customKey.Pods[0], anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}
	return nil
}

func ShareContent(user response.User, anotherUserEmail string, allKeys AllKeys) error {
	err := shareCustomBlockAndPod(user, anotherUserEmail, allKeys.CustomKey)
	if err != nil {
		return err
	}

	err = shareClassroomContent(user, anotherUserEmail, allKeys.TeacherKey)
	if err != nil {
		return err
	}

	err = shareProjectManagementContent(user, anotherUserEmail, allKeys.ProjectKey)
	if err != nil {
		return err
	}
	return nil
}
