package recipes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/endpoints/conversations"
	"development/go/recipes/lib/endpoints/profile"
	"development/go/recipes/lib/endpoints/registration"
	"development/go/recipes/lib/structs/request"

	recipes2 "development/go/recipes/lib/helpers/recipes"
	response2 "development/go/recipes/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

// CreatePrivateConversation Sign up, activate user, sign in, get all keys.
func CreatePrivateConversation() {
	log.Info("Objective: Send messages to a private conversation")
	_, err := recipes2.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes2.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	var anotherUser response2.User
	var anotherProfile response2.Profile
	anotherUser, anotherProfile, err = getUserProfile()
	if err != nil {
		return
	}

	var conversation response2.Conversation
	conversation, err = createConversation(anotherProfile, user)
	if err != nil {
		return
	}

	err = displayMessages(conversation, anotherUser)
	if err != nil {
		return
	}
}

func displayMessages(conversation response2.Conversation, anotherUser response2.User) error {
	log.Info("Retrieve messages as another user")
	recipes2.SleepBefore()
	var err error
	conversation, err = conversations.GetConversation(anotherUser.JwtToken, conversation.ID)
	if err != nil {
		return err
	}
	for index, message := range *conversation.Messages {
		log.Printf(".Message %d: %s (sent by %s at %s)", index, message.MessageText, message.AddedBy, message.MessageTime)
	}
	recipes2.SleepAfter()
	return nil
}

func createConversation(anotherProfile response2.Profile, user response2.User) (response2.Conversation, error) {
	log.Info("Create a private conversation")
	recipes2.SleepBefore()
	conversationBody := conversations.ConversationReqBody{
		MessageText: "hey there!",
		Usernames:   anotherProfile.Username,
	}
	conversation, err := conversations.AddPrivateOrGroupConversation(user.JwtToken, conversationBody)
	if err != nil {
		return response2.Conversation{}, err
	}
	log.Info(".Conversation created")
	recipes2.SleepAfter()
	return conversation, nil
}

func getUserProfile() (response2.User, response2.Profile, error) {
	log.Info("Get profiles of (target) users to create conversation")
	recipes2.SleepBefore()
	anotherUser, err := registration.SignInByEmail(request.SignInReqBody{
		Email:    lib.DefaultEmail,
		Password: lib.Password,
	})
	if err != nil {
		return response2.User{}, response2.Profile{}, err
	}
	var anotherProfile response2.Profile
	anotherProfile, err = profiles.GetUsersProfile(anotherUser.JwtToken)
	if err != nil {
		return response2.User{}, response2.Profile{}, err
	}
	return anotherUser, anotherProfile, nil
}
