package dashboard_1

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetDashboardDetails(jwtToken string) (response.Dashboard, error) {
	resDashboard := response.Dashboard{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteDashboardGetDashboardDetails)
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
	res, err = client.Do(req)
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
