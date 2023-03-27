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

func GetScalesAlongWithBlocksAndPodsBasedOnThem(jwtToken string, keyId string) (response.ScalesKey, error) {
	resScalesKey := response.ScalesKey{}
	route, err := helpers2.GetRoute(lib.RouteKeysGetScalesAlongWithBlocksAndPodsBasedOnThem, keyId)
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resScalesKey, err
	}

	defer helpers2.CloseBody(res.Body)

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
