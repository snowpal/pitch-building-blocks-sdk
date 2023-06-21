package attributes

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetResourceAttrs(jwtToken string) (response.ResourceAttributes, error) {
	var resAttributes response.ResourceAttributes
	route, err := helpers.GetRoute(lib.RouteAttributesGetDisplayableAttributes)
	if err != nil {
		return resAttributes, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resAttributes, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		return resAttributes, err
	}

	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		return resAttributes, err
	}

	err = json.Unmarshal(body, &resAttributes)
	if err != nil {
		return resAttributes, err
	}

	return resAttributes, nil
}
