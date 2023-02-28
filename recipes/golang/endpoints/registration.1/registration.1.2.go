package registration_1

import (
	"development/go/recipies/structs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SignIn() (structs.User, error) {
	var userRegistration structs.User

	url := "https://gateway-dev.snowpal.com/app/users/sign-in"
	method := "POST"

	payload := strings.NewReader(`{
		"email": "apiuser3@gmail.com",
		"password": "Welcome1!"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return userRegistration, err
	}
	req.Header.Add("x-api-key", "wf8sHELzWp9MGizZME5Zsjk4IntZS0e8mdYMYjjg")
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

	body, err := io.ReadAll(res.Body)
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
