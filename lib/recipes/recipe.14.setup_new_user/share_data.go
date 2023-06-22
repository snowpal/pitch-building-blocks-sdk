package setupnewuser

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func ShareContent(user response.User, anotherUserEmail string, allKeys AllKeys) error {
	log.Info("Share custom key data with ", anotherUserEmail)
	err := shareCustomKeyData(user, anotherUserEmail, allKeys.CustomKey)
	if err != nil {
		return err
	}

	log.Info("Share teacher key data with ", anotherUserEmail)
	err = shareClassroomKeyData(user, anotherUserEmail, allKeys.TeacherKey)
	if err != nil {
		return err
	}

	log.Info("Share project key data with ", anotherUserEmail)
	err = shareProjectKeyData(user, anotherUserEmail, allKeys.ProjectKey)
	if err != nil {
		return err
	}
	return nil
}

func shareCustomKeyData(user response.User, anotherUserEmail string, customKey KeyWithResources) error {
	log.Info("Share block ", customKey.firstBlock().Block.Name, " for custom key, ", customKey.Key.Name)
	err := recipes.SearchUserAndShareBlock(user, customKey.firstBlock().Block, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}

	log.Info("Share block pod ", customKey.firstBlockPod().Name, " for custom key, ", customKey.Key.Name, " and for ", customKey.firstBlock().Block.Name)
	err = recipes.SearchUserAndShareBlockPod(user, customKey.firstBlockPod(), anotherUserEmail, lib.WriteAcl)
	if err != nil {
		return err
	}

	log.Info("Share key pod ", customKey.firstPod().Name, " for custom key, ", customKey.Key.Name)
	err = recipes.SearchUserAndSharePod(user, customKey.firstPod(), anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}
	return nil
}

func shareClassroomKeyData(user response.User, anotherUserEmail string, teacherKey KeyWithResources) error {
	log.Info("Share block ", teacherKey.firstBlock().Block.Name, " for teacher key, ", teacherKey.Key.Name)
	err := recipes.SearchUserAndShareBlock(user, teacherKey.firstBlock().Block, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}
	return nil
}

func shareProjectKeyData(user response.User, anotherUserEmail string, projectKey KeyWithResources) error {
	log.Info("Share block ", projectKey.firstBlock().Block.Name, " for project key, ", projectKey.Key.Name)
	err := recipes.SearchUserAndShareBlock(user, projectKey.firstBlock().Block, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}
	return nil
}
