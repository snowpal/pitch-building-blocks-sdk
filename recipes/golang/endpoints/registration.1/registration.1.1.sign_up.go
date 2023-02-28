package registration_1

import (
	"development/go/recipes/endpoints"
	"development/go/recipes/structs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Signup2(email string) (structs.UserSignUp, error) {
	var userSignUp structs.UserSignUp
	payload := strings.NewReader(fmt.Sprintf(`{
		"email": "%s",
		"password": "Welcome1!",
		"confirmPassword": "Welcome1!"
	}`, email))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, endpoints.UrlSignUp, payload)

	if err != nil {
		fmt.Println(err)
		return userSignUp, err
	}
	req.Header.Add("x-api-key", endpoints.ApiKey)
	req.Header.Add("Content-Type", "application/json")

	fmt.Println(req.Body)
	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return userSignUp, err
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
		return userSignUp, err
	}

	err = json.Unmarshal(body, &userSignUp)
	if err != nil {
		return userSignUp, err
	}

	fmt.Println("userSignUp", userSignUp)
	return userSignUp, err
}

func Signup(email string) (structs.UserSignedUp, error) {
	var userSignedUp structs.UserSignedUp
	payload := strings.NewReader(fmt.Sprintf(`{
		"email": "%s",
		"password": "Welcome1!",
		"confirmPassword": "Welcome1!"
	}`, email))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, endpoints.UrlSignUp, payload)

	if err != nil {
		fmt.Println(err)
		return userSignedUp, err
	}
	req.Header.Add("x-api-key", endpoints.ApiKey)
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

	fmt.Println(string(body))
	err = json.Unmarshal(body, &userSignedUp)
	if err != nil {
		return userSignedUp, err
	}

	return userSignedUp, nil
}
