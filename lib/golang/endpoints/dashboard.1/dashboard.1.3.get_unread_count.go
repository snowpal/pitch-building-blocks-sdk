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

func GetUnreadCount(jwtToken string) (response.DashboardUnreadCount, error) {
	dashboardUnreadCount := response.DashboardUnreadCount{}
	route, err := helpers.GetRoute(golang.RouteDashboardGetUnreadCount)
	if err != nil {
		fmt.Println(err)
		return dashboardUnreadCount, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return dashboardUnreadCount, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return dashboardUnreadCount, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return dashboardUnreadCount, err
	}

	err = json.Unmarshal(body, &dashboardUnreadCount)
	if err != nil {
		fmt.Println(err)
		return dashboardUnreadCount, err
	}
	return dashboardUnreadCount, nil
}
