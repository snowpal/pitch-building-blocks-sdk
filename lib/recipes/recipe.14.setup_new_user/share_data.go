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
	log.Info(fmt.Sprintf("Share block %s for custom key, %s", customKey.firstBlock().Block.Name, customKey.Key.Name))
	err := recipes.SearchUserAndShareBlock(user, customKey.firstBlock().Block, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Share block pod %s for custom key, %s  and for %s block", customKey.firstBlockPod().Name, customKey.Key.Name, customKey.firstBlock().Block.Name))
	err = recipes.SearchUserAndShareBlockPod(user, customKey.firstBlockPod(), anotherUserEmail, lib.WriteAcl)
	if err != nil {
		return err
	}

	log.Info(fmt.Sprintf("Share key pod %s for custom key, %s", customKey.firstPod().Name, customKey.Key.Name))
	err = recipes.SearchUserAndSharePod(user, customKey.firstPod(), anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}
	return nil
}

func shareClassroomKeyData(user response.User, anotherUserEmail string, teacherKey KeyWithResources) error {
	log.Info(fmt.Sprintf("Share block %s for teacher key, %s", teacherKey.firstBlock().Block.Name, teacherKey.Key.Name))
	err := recipes.SearchUserAndShareBlock(user, teacherKey.firstBlock().Block, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}
	return nil
}

func shareProjectKeyData(user response.User, anotherUserEmail string, projectKey KeyWithResources) error {
	log.Info(fmt.Sprintf("Share block %s for project key, %s", projectKey.firstBlock().Block.Name, projectKey.Key.Name))
	err := recipes.SearchUserAndShareBlock(user, projectKey.firstBlock().Block, anotherUserEmail, lib.ReadAcl)
	if err != nil {
		return err
	}
	return nil
}
