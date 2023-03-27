package keys

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlockTypesAndBlocksBasedOnThemInKey(jwtToken string, keyId string) (response.BlockTypesKey, error) {
	resBlockTypesKey := response.BlockTypesKey{}
	route, err := helpers2.GetRoute(lib.RouteKeysGetBlockTypesAndBlocksBasedOnThemInKey, keyId)
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockTypesKey, err
	}

	defer helpers2.CloseBody(res.Body)

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
