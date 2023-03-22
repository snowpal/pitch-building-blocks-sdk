package main

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers/recipes"
	"development/go/recipes/lib/golang/structs/response"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Objective: Assign and Publish Pod Grades for a Student")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(golang.DefaultEmail, golang.Password)
	if err != nil {
		return
	}

	key, err := addKey(user, "")
	_, err = addBlock(user, "", key)
	pod, err := addPod(user, "", key)
	assignPodGrade(user, pod)
}

func assignPodGrade(user response.User, pod response.Pod) {
}
