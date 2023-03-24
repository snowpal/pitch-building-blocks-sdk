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

func GetLatestVersion() (response.Version, error) {
	resVersion := response.Version{}
	route, err := helpers.GetRoute(building_blocks.RouteVersionGetLatestVersion)
	if err != nil {
		fmt.Println(err)
		return resVersion, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resVersion, err
	}

	helpers.AddAppHeaders(req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resVersion, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resVersion, err
	}

	err = json.Unmarshal(body, &resVersion)
	if err != nil {
		fmt.Println(err)
		return resVersion, err
	}
	return resVersion, nil
}
