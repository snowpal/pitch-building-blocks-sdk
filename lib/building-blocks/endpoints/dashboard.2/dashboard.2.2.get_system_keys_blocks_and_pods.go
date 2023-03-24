package dashboard_2

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetSystemKeysBlocksAndPods(jwtToken string) ([]response.UserKey, error) {
	resUserKeys := response.UserKeys{}
	route, err := helpers.GetRoute(building_blocks.RouteDashboardGetSystemKeysBlocksAndPods)
	if err != nil {
		fmt.Println(err)
		return resUserKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resUserKeys.Keys, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resUserKeys.Keys, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resUserKeys.Keys, err
	}

	err = json.Unmarshal(body, &resUserKeys)
	if err != nil {
		fmt.Println(err)
		return resUserKeys.Keys, err
	}
	return resUserKeys.Keys, nil
}