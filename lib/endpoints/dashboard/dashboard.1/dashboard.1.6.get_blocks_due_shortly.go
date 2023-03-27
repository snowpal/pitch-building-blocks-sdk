package dashboard

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetBlocksDueShortly(jwtToken string) (response.Dashboard, error) {
	resDashboard := response.Dashboard{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetBlocksDueShortly)
	if err != nil {
		fmt.Println(err)
		return resDashboard, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resDashboard, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resDashboard, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resDashboard, err
	}

	err = json.Unmarshal(body, &resDashboard)
	if err != nil {
		fmt.Println(err)
		return resDashboard, err
	}
	return resDashboard, nil
}
