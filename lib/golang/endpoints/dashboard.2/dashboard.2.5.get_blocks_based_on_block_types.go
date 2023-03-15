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

func GetBlocksBasedOnBlockTypes(jwtToken string) ([]response.BlockTypesKey, error) {
	resBlockTypesKeys := response.BlockTypesKeys{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteDashboardGetBlocksBasedOnBlockTypes)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKeys.Keys, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKeys.Keys, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKeys.Keys, err
	}

	err = json.Unmarshal(body, &resBlockTypesKeys)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKeys.Keys, err
	}
	return resBlockTypesKeys.Keys, nil
}
