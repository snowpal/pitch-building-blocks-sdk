package registration

import (
	go_status "development/go/recipes/lib/go-status"
	"development/go/recipes/lib/go-status/helpers"
	"development/go/recipes/lib/go-status/structs/request"
	"fmt"
	"net/http"
	"strings"
)

func ResetPassword(jwtToken string, reqBody request.ResetPasswordReqBody) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(go_status.RouteRegistrationResetPassword)
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
