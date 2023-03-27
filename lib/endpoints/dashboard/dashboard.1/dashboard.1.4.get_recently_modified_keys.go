package dashboard

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetRecentlyModifiedKeys(jwtToken string) ([]common.SlimKey, error) {
	resRecentKeys := response.RecentlyModifiedKeys{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetRecentlyModifiedKeys)
	if err != nil {
		fmt.Println(err)
		return resRecentKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resRecentKeys.Keys, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resRecentKeys.Keys, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resRecentKeys.Keys, err
	}

	err = json.Unmarshal(body, &resRecentKeys)
	if err != nil {
		fmt.Println(err)
		return resRecentKeys.Keys, err
	}
	return resRecentKeys.Keys, nil
}
