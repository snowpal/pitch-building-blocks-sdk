package registration_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/endpoints/structs"

	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SignIn(email string) (structs.UserSignedIn, error) {
	var userRegistration structs.UserSignedIn
	payload := strings.NewReader(fmt.Sprintf(`{
		"email": "%s",
		"password": "Welcome1!"
	}`, email))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, golang.UrlSignIn, payload)

	if err != nil {
		fmt.Println(err)
		return userRegistration, err
	}
	req.Header.Add("x-api-key", golang.XApiKey)
	req.Header.Add("Content-Type", "application/json")

	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return userRegistration, err
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
		return userRegistration, err
	}

	err = json.Unmarshal(body, &userRegistration)
	if err != nil {
		return userRegistration, err
	}

	return userRegistration, nil
}
