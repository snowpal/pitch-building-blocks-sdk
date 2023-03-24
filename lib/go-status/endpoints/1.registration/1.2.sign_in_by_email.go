package registration

import (
	go_status "development/go/recipes/lib/go-status"
	"development/go/recipes/lib/go-status/helpers"
	"development/go/recipes/lib/go-status/structs/request"
	"development/go/recipes/lib/go-status/structs/response"
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
	var route string
	route, err = helpers.GetRoute(go_status.RouteRegistrationSignInByEmail)
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

	res, err := helpers.MakeRequest(req)
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
