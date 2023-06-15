package recipes

import "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

func ShareProjectManagementContent(user response.User, anotherUserEmail string, projectKey KeyWithResources) error {
	var err error
	return err
}

func ShareClassroomContent(user response.User, anotherUserEmail string, teacherKey KeyWithResources) error {
	var err error
	return err
}

func ShareCustomManagementContent(user response.User, anotherUserEmail string, customKey KeyWithResources) error {
	var err error
	return err
}

func ShareContent(user response.User, anotherUserEmail string, allKeys AllKeys) error {
	var err error
	err = ShareCustomManagementContent(user, anotherUserEmail, allKeys.CustomKey)
	if err != nil {
		return err
	}

	err = ShareClassroomContent(user, anotherUserEmail, allKeys.TeacherKey)
	if err != nil {
		return err
	}

	err = ShareProjectManagementContent(user, anotherUserEmail, allKeys.ProjectKey)
	if err != nil {
		return err
	}

	return nil
}
