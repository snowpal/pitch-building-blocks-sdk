package attributes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	helpers2 "development/go/recipes/lib/helpers"
)

func GetResourceAttrs(jwtToken string) (response.ResourceAttributes, error) {
	var resAttributes response.ResourceAttributes
	route, err := helpers2.GetRoute(lib.RouteAttributesGetDisplayableAttributes)
	if err != nil {
		fmt.Println(err)
		return resAttributes, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAttributes, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAttributes, err
	}

	defer helpers2.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAttributes, err
	}

	err = json.Unmarshal(body, &resAttributes)
	if err != nil {
		fmt.Println(err)
		return resAttributes, err
	}

	return resAttributes, nil
}
