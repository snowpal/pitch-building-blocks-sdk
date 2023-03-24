package attributes_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetResourceAttrs(jwtToken string) (response.ResourceAttributes, error) {
	var resAttributes response.ResourceAttributes
	route, err := helpers.GetRoute(golang.RouteAttributesGetDisplayableAttributes)
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

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAttributes, err
	}

	defer helpers.CloseBody(res.Body)

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
