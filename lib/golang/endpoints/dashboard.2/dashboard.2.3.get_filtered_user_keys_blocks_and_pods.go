package dashboard_2

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetFilteredUserKeysBlocksAndPods(jwtToken string) ([]response.FilteredKey, error) {
	resFilteredUserKeys := response.FilteredKeys{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteDashboardGetFilteredUserKeysBlocksAndPods)
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

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resFilteredUserKeys.Keys, err
	}

	defer helpers.CloseBody(res.Body)

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
