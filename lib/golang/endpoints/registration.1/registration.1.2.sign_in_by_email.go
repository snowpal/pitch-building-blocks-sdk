package registration_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"

	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SignIn(email string) (response.UserRegistration, error) {
	var userRegistration response.UserRegistration
	fmt.Println("TODO: Replace with struct")
	payload := strings.NewReader(fmt.Sprintf(`{
		"email": "%s",
		"password": "Welcome1!"
	}`, email))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, helpers.GetRoute(golang.RouteRegistrationSignInByEmail), payload)

	if err != nil {
		fmt.Println(err)
		return userRegistration, err
	}

	helpers.AddAppHeaders(req)
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
