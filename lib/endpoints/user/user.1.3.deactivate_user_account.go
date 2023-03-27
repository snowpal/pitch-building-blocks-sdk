package user

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"fmt"
	"net/http"
)

func DeactivateUserAccount(jwtToken string) error {
	route, err := helpers.GetRoute(lib.RouteUserDeactivateUserAccount)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
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
