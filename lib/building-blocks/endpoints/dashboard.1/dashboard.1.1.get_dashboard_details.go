package dashboard_1

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetDashboardDetails(jwtToken string) (response.Dashboard, error) {
	resDashboard := response.Dashboard{}
	route, err := helpers.GetRoute(building_blocks.RouteDashboardGetDashboardDetails)
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

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resDashboard, err
	}

	defer helpers.CloseBody(res.Body)

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