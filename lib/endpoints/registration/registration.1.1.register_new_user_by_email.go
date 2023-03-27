package registration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func RegisterNewUserByEmail(reqBody request.SignupReqBody) (response.User, error) {
	resUserRegistration := response.UserRegistration{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resUserRegistration.User, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(lib.RouteRegistrationRegisterNewUserByEmail)
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
