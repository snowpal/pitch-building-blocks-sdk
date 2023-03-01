package registration_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/structs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Signup(email string) (structs.UserSignedUp, error) {
	var userSignedUp structs.UserSignedUp
	fmt.Println("TODO: Replace with struct")
	payload := strings.NewReader(fmt.Sprintf(`{
		"email": "%s",
		"password": "Welcome1!",
		"confirmPassword": "Welcome1!"
	}`, email))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, golang.UrlSignUp, payload)

	if err != nil {
		fmt.Println(err)
		return userSignedUp, err
	}
	req.Header.Add("x-api-key", golang.XApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return userSignedUp, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return userSignedUp, err
	}

	err = json.Unmarshal(body, &userSignedUp)
	if err != nil {
		return userSignedUp, err
	}

	return userSignedUp, nil
}
