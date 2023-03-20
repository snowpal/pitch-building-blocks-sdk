package registration_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"

	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SignInByEmail(reqBody request.SignInReqBody) (response.User, error) {
	resUserRegistration := response.UserRegistration{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resUserRegistration.User, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteRegistrationSignInByEmail)
	if err != nil {
		fmt.Println(err)
		return resUserRegistration.User, err
	}
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resUserRegistration.User, err
	}

	helpers.AddAppHeaders(req)

	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resUserRegistration.User, err
	}

	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resUserRegistration.User, err
	}

	err = json.Unmarshal(body, &resUserRegistration)
	if err != nil {
		fmt.Println(err)
		return resUserRegistration.User, err
	}
	return resUserRegistration.User, nil
}
