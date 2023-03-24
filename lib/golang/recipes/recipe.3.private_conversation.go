package main

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/endpoints/conversations.1"
	profiles "development/go/recipes/lib/golang/endpoints/profile.1"
	registration "development/go/recipes/lib/golang/endpoints/registration.1"
	"development/go/recipes/lib/golang/helpers/recipes"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	log "github.com/sirupsen/logrus"
)

const Password string = "Welcome1!"
const (
	User1Email string = "mike@yopmail.com"
)

// Sign up, activate user, sign in, get all keys.
func CreatePrivateConversation() {
	log.Info("Objective: Send messages to a private conversation")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(golang.ActiveUser, golang.Password)
	if err != nil {
		return
	}

	var anotherUser response.User
	var anotherProfile response.Profile
	anotherUser, anotherProfile, err = getUserProfile()
	if err != nil {
		return
	}

	var conversation response.Conversation
	conversation, err = createConversation(anotherProfile, user)
	if err != nil {
		return
	}

	err = displayMessages(conversation, anotherUser)
	if err != nil {
		return
	}
}

func displayMessages(conversation response.Conversation, anotherUser response.User) error {
	log.Info("Retrieve messages as another user")
	recipes.SleepBefore()
	var err error
	conversation, err = conversations.GetConversation(anotherUser.JwtToken, conversation.ID)
	if err != nil {
		return err
	}
	for index, message := range *conversation.Messages {
		log.Printf(".Message %d: %s (sent by %s at %s)", index, message.MessageText, message.AddedBy, message.MessageTime)
	}
	recipes.SleepAfter()
	return nil
}

func createConversation(anotherProfile response.Profile, user response.User) (response.Conversation, error) {
	log.Info("Create a private conversation")
	recipes.SleepBefore()
	conversationBody := conversations.ConversationReqBody{
		MessageText: "hey there!",
		Usernames: []string{
			anotherProfile.Username,
		},
	}
	conversation, err := conversations.AddPrivateOrGroupConversation(user.JwtToken, conversationBody)
	if err != nil {
		return response.Conversation{}, err
	}
	log.Info(".Conversation created")
	recipes.SleepAfter()
	return conversation, nil
}

func getUserProfile() (response.User, response.Profile, error) {
	log.Info("Get profiles of (target) users to create conversation")
	recipes.SleepBefore()
	anotherUser, err := registration.SignInByEmail(request.SignInReqBody{
		Email:    User1Email,
		Password: Password,
	})
	if err != nil {
		return response.User{}, response.Profile{}, err
	}
	var anotherProfile response.Profile
	anotherProfile, err = profiles.GetUsersProfile(anotherUser.JwtToken)
	if err != nil {
		return response.User{}, response.Profile{}, err
	}
	return anotherUser, anotherProfile, nil
}
