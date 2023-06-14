package recipes

import (
	log "github.com/sirupsen/logrus"
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func displayUser(email string) {
	user, err := recipes.SignIn(email, lib.Password)
	if err != nil {
		return
	}
	log.Info("- %s | %s", email, user.JwtToken)
}

func displayAllKeys(user response.User) {
	// Fetch all keys
	// Iterate through keys
	// Diplay key
	// Fetch all blocks
	// Iterate through blocks
	// Display block
	// Fetch all block pods
	// Iterate through block pods
	// Display block pod
	// Fetch all key pods
	// Iterate through key pods
	// Display key pod
}

func displayAllNotifications(anotherUser response.User) {
	// Fetch all notifications
	// Iterate through notifications
	// Display notification
}

func DisplayContent(user response.User, anotherUserEmail string) {
	log.Info("## Registered Users")
	displayUser(user.Email)
	displayUser(anotherUserEmail)

	log.Info("## Resources Created for user: %s", user.Email)
	displayAllKeys(user)

	anotherUser, err := recipes.SignIn(anotherUserEmail, lib.Password)
	if err != nil {
		return
	}

	log.Info("## Notifications for shared content as user: %s", anotherUserEmail)
	displayAllNotifications(anotherUser)
}
