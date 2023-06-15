package recipes

import (
	"fmt"

	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GenerateDynamicEmail(nextNumber int) string {
	return fmt.Sprintf("apiuser_rec%d_bb", nextNumber)
}

func RegisterNewUser() (string, error) {
	var err error
	var user response.User
	for i := 1; i <= 1000; i++ {
		user, err = recipes.RegisterUser(GenerateDynamicEmail(i))
		if !user.Inactive {
			println(user.Email)
			break
		} else if err != nil {
			return user.Email, err
		}
	}
	return user.Email, nil
}
