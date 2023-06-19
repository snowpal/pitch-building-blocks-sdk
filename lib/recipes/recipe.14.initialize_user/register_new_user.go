package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

func GenerateDynamicEmail(nextNumber int) string {
	return fmt.Sprintf("apiuser_rec%d_bb@yopmail.com", nextNumber)
}

func RegisterNewUser() (string, error) {
	var err error
	var user response.User

	for i := 1; ; i += 1 {
		email := GenerateDynamicEmail(i)
		log.Info("Register new user with ", email)
		user, err = recipes.RegisterUser(email)
		if err != nil {
			log.Info(email, " is already registered.")
		} else {
			return user.Email, nil
		}
	}
}
