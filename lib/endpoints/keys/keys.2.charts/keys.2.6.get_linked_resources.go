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

func GetLinkedResources(jwtToken string, keyId string) (response.LinkedResources, error) {
	resLinkedResources := response.LinkedResources{}
	route, err := helpers2.GetRoute(lib.RouteKeysGetLinkedResources, keyId)
	if err != nil {
		fmt.Println(err)
		return resLinkedResources, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resLinkedResources, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resLinkedResources, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resLinkedResources, err
	}

	err = json.Unmarshal(body, &resLinkedResources)
	if err != nil {
		fmt.Println(err)
		return resLinkedResources, err
	}
	return resLinkedResources, nil
}
