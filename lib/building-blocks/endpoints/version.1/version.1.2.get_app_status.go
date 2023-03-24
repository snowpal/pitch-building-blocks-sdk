package version

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAppStatus() (response.Version, error) {
	resAppStatus := response.Version{}
	route, err := helpers.GetRoute(building_blocks.RouteVersionGetAppStatus)
	if err != nil {
		fmt.Println(err)
		return resAppStatus, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAppStatus, err
	}

	helpers.AddAppHeaders(req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAppStatus, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAppStatus, err
	}

	err = json.Unmarshal(body, &resAppStatus)
	if err != nil {
		fmt.Println(err)
		return resAppStatus, err
	}
	return resAppStatus, nil
}