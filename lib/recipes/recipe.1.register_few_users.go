package recipes

import (
	log "github.com/sirupsen/logrus"
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
)

func RegisterFewUsers() {
	log.Info("Objective: Sign up a few new users")
	_, err := recipes.RegisterUser(lib.DefaultEmail)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(lib.ActiveUser)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(lib.AdminUser)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(lib.ReadUser)
	if err != nil {
		return
	}

	_, err = recipes.RegisterUser(lib.WriteUser)
	if err != nil {
		return
	}
}
