package recipes

import "github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

func createProjectManagementContent(user response.User) (KeyWithContent, error) {
	var keyWithContent KeyWithContent
	var err error

	return keyWithContent, err
}

func createClassroomContent(user response.User) (KeyWithContent, error) {
	var keyWithContent KeyWithContent
	var err error

	return keyWithContent, err
}

func createCustomManagementContent(user response.User) (KeyWithContent, error) {
	var keyWithContent KeyWithContent
	var err error

	return keyWithContent, err
}

func CreateContent(user response.User) (AllKeys, error) {
	var allKeys AllKeys
	var err error

	var customKey KeyWithContent
	customKey, err = createCustomManagementContent(user)
	if err != nil {
		return allKeys, err
	}
	allKeys.CustomKey = customKey

	var classroomKey KeyWithContent
	classroomKey, err = createClassroomContent(user)
	if err != nil {
		return allKeys, err
	}
	allKeys.ClassroomKey = classroomKey

	var projectKey KeyWithContent
	projectKey, err = createProjectManagementContent(user)
	if err != nil {
		return allKeys, err
	}
	allKeys.ProjectKey = projectKey

	return allKeys, err
}
