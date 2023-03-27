package registration

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func SignInByEmail(reqBody request.SignInReqBody) (response.User, error) {
	resUserRegistration := response.UserRegistration{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resUserRegistration.User, err
	}
	payload := strings.NewReader(requestBody)
	var route string
	route, err = helpers2.GetRoute(lib.RouteRegistrationSignInByEmail)
	if err != nil {
		fmt.Println(err)
		return resUserRegistration.User, err
	}
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resUserRegistration.User, err
	}

	helpers2.AddAppHeaders(req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resUserRegistration.User, err
	}

	defer helpers2.CloseBody(res.Body)

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
