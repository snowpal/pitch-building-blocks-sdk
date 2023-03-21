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

func GetBlocksAndPodsBasedOnScales(jwtToken string) ([]response.ScalesKey, error) {
	resScalesKeys := response.ScalesKeys{}
	route, err := helpers.GetRoute(golang.RouteDashboardGetBlocksAndPodsBasedOnScales)
	if err != nil {
		fmt.Println(err)
		return *resScalesKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return *resScalesKeys.Keys, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return *resScalesKeys.Keys, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return *resScalesKeys.Keys, err
	}

	err = json.Unmarshal(body, &resScalesKeys)
	if err != nil {
		fmt.Println(err)
		return *resScalesKeys.Keys, err
	}
	return *resScalesKeys.Keys, nil
}
