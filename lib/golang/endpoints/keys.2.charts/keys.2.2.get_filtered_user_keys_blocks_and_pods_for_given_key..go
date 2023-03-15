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

func GetFilteredUserKeysBlocksAndPodsForGivenKey(jwtToken string, keyId string) (response.FilteredKey, error) {
	resFilteredUserKey := response.FilteredKey{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteKeysGetFilteredUserKeysBlocksAndPodsForGivenKey, keyId)
	if err != nil {
		fmt.Println(err)
		return resFilteredUserKey, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resFilteredUserKey, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resFilteredUserKey, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resFilteredUserKey, err
	}

	err = json.Unmarshal(body, &resFilteredUserKey)
	if err != nil {
		fmt.Println(err)
		return resFilteredUserKey, err
	}
	return resFilteredUserKey, nil
}
