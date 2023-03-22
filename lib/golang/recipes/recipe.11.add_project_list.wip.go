package main

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers/recipes"
	"development/go/recipes/lib/golang/structs/response"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Objective: Add Project Lists, Project Pods, Move Pod between Lists")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(golang.DefaultEmail, golang.Password)
	if err != nil {
		return
	}

	key, err := addKey(user, "")
	block, err := addBlock(user, "", key)
	pod, err := addBlockPod(user, "", key, block)
	projectList1, err := addProjectList(user, "", key, block)
	projectList2, err := addProjectList(user, "", key, block)
	movePodBetweenLists(projectList1, projectList2, pod)
}

func addBlockPod(user response.User, podName string, key response.Key, block response.Block) (response.Pod, error) {
	return response.Pod{}, nil
}

func movePodBetweenLists(projectList1 response.ProjectList, projectList2 response.ProjectList, pod response.Pod) {
}

func addProjectList(user response.User, projectListName string, key response.Key,
	block response.Block) (response.ProjectList, error) {
	return response.ProjectList{}, nil
}
