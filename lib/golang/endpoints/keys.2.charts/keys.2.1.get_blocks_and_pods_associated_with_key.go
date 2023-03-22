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

func GetBlocksAndPodsAssociatedWithKey(jwtToken string, keyId string) (response.UserKey, error) {
	resUserKey := response.UserKey{}
	route, err := helpers.GetRoute(golang.RouteKeysGetBlocksAndPodsAssociatedWithKey, keyId)
	if err != nil {
		fmt.Println(err)
		return resUserKey, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resUserKey, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resUserKey, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resUserKey, err
	}

	err = json.Unmarshal(body, &resUserKey)
	if err != nil {
		fmt.Println(err)
		return resUserKey, err
	}
	return resUserKey, nil
}
