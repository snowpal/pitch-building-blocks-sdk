package dashboard

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetFilteredUserKeysBlocksAndPods(jwtToken string) ([]response.FilteredKey, error) {
	resFilteredUserKeys := response.FilteredKeys{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetFilteredUserKeysBlocksAndPods)
	if err != nil {
		fmt.Println(err)
		return resFilteredUserKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resFilteredUserKeys.Keys, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resFilteredUserKeys.Keys, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resFilteredUserKeys.Keys, err
	}

	err = json.Unmarshal(body, &resFilteredUserKeys)
	if err != nil {
		fmt.Println(err)
		return resFilteredUserKeys.Keys, err
	}
	return resFilteredUserKeys.Keys, nil
}
