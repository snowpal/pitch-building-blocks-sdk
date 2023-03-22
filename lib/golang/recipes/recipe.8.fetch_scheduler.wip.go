package main

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers/recipes"
	"development/go/recipes/lib/golang/structs/response"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Objective: Set Due Date for Block and Pod, and Fetch Scheduler Events")
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
	_, err = addPod(user, "", key)
	fetchSchedulerEvents()
}

func fetchSchedulerEvents() {
}

func addPod(user response.User, podName string, key response.Key) (response.Pod, error) {
	return response.Pod{}, nil
}

func addKey(user response.User, keyName string) (response.Key, error) {
	return response.Key{}, nil
}
