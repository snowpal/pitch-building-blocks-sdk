package main

import (
	keys "development/go/recipies/endpoints/keys.1"
	registration "development/go/recipies/endpoints/registration.1"
)

func main() {
	email := "api_user_1@yopmail.com"
	registration.Signup(email)
	registration.Activate()
	registration.SignIn()
	keys.GetAllKeys()
}
