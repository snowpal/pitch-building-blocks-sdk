package keys_2

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlockTypesAndBlocksBasedOnThemInKey(jwtToken string, keyId string) (response.BlockTypesKey, error) {
	resBlockTypesKey := response.BlockTypesKey{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteKeysGetBlockTypesAndBlocksBasedOnThemInKey, keyId)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKey, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKey, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKey, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKey, err
	}

	err = json.Unmarshal(body, &resBlockTypesKey)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKey, err
	}
	return resBlockTypesKey, nil
}
