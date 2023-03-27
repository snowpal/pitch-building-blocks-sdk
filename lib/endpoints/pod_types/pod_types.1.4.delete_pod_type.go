package podTypes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"fmt"
	"net/http"
)

func DeletePodType(jwtToken string, podTypeId string) error {
	route, err := helpers.GetRoute(lib.RoutePodTypesDeletePodType, podTypeId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, route, nil)
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
