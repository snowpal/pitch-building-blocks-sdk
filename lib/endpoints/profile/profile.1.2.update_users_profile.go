package profiles

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"fmt"
	"net/http"
	"strings"
)

type ProfileReqBody struct {
	FirstName   *string `json:"firstName"`
	MiddleName  *string `json:"middleName"`
	LastName    *string `json:"lastName"`
	PhoneNumber *string `json:"phoneNumber"`
}

func UpdateUsersProfile(jwtToken string, reqBody ProfileReqBody) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(lib.RouteProfileUpdateUsersProfile)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
