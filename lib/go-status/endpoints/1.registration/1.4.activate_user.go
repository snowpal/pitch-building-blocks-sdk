package registration

import (
	go_status "development/go/recipes/lib/go-status"
	"development/go/recipes/lib/go-status/helpers"
	"fmt"
	"net/http"
)

func ActivateUser(userId string) error {
	route, err := helpers.GetRoute(go_status.RouteRegistrationActivateUser, userId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddAppHeaders(req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
