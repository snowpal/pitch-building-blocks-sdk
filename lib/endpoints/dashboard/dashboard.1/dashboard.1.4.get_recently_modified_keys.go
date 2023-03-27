package dashboard

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetRecentlyModifiedKeys(jwtToken string) ([]common.SlimKey, error) {
	resRecentKeys := response.RecentlyModifiedKeys{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetRecentlyModifiedKeys)
	if err != nil {
		fmt.Println(err)
		return resRecentKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resRecentKeys.Keys, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resRecentKeys.Keys, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resRecentKeys.Keys, err
	}

	err = json.Unmarshal(body, &resRecentKeys)
	if err != nil {
		fmt.Println(err)
		return resRecentKeys.Keys, err
	}
	return resRecentKeys.Keys, nil
}
