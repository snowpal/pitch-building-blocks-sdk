package main

import (
	"development/go/recipes/lib/golang"
	registration "development/go/recipes/lib/golang/endpoints/registration.1"
	log "github.com/sirupsen/logrus"

	"fmt"
)

func main() {
	log.Info(".sign in user email: ", golang.Email)
	_, err := registration.SignIn(golang.Email)
	if err != nil {
		fmt.Println(err)
	}
}
