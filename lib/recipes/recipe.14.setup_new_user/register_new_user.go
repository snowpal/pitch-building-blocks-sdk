package setupnewuser

import (
	"fmt"

	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
)

// RegisterNewUser Register a new user. Ignore errors due to existing users, keep trying till we get successful.
func RegisterNewUser() (string, error) {
	var err error
	var user response.User
	for i := 1; ; i += 1 {
		email := fmt.Sprintf("apiuser_rec%d_bb@yopmail.com", i)
		log.Info("Register new user with ", email)
		user, err = recipes.RegisterUser(email)
		if err != nil {
			log.Info(email, " is already registered.")
		} else {
			return user.Email, nil
		}
	}
}
