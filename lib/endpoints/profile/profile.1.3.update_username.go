package profiles

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"fmt"
	"net/http"
)

func UpdateUsername(jwtToken string, username string) error {
	route, err := helpers.GetRoute(lib.RouteProfileUpdateUsername, username)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, nil)
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
