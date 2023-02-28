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

func SignIn(email string) (structs.User, error) {
	var userRegistration structs.User
	payload := strings.NewReader(fmt.Sprintf(`{
		"email": "%s",
		"password": "Welcome1!"
	}`, email))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, endpoints.Url, payload)

	if err != nil {
		fmt.Println(err)
		return userRegistration, err
	}
	req.Header.Add("x-api-key", endpoints.ApiKey)
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
