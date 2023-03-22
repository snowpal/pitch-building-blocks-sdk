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

func GetScalesAlongWithBlocksAndPodsBasedOnThem(jwtToken string, keyId string) (response.ScalesKey, error) {
	resScalesKey := response.ScalesKey{}
	route, err := helpers.GetRoute(golang.RouteKeysGetScalesAlongWithBlocksAndPodsBasedOnThem, keyId)
	if err != nil {
		fmt.Println(err)
		return resScalesKey, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resScalesKey, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resScalesKey, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resScalesKey, err
	}

	err = json.Unmarshal(body, &resScalesKey)
	if err != nil {
		fmt.Println(err)
		return resScalesKey, err
	}
	return resScalesKey, nil
}
