package dashboard_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetRecentlyModifiedKeys(jwtToken string) ([]common.SlimKey, error) {
	resRecentKeys := response.RecentlyModifiedKeys{}
	route, err := helpers.GetRoute(golang.RouteDashboardGetRecentlyModifiedKeys)
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

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resRecentKeys.Keys, err
	}

	defer helpers.CloseBody(res.Body)

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
