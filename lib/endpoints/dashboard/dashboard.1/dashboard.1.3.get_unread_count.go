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

func GetUnreadCount(jwtToken string) (response.DashboardUnreadCount, error) {
	dashboardUnreadCount := response.DashboardUnreadCount{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetUnreadCount)
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return dashboardUnreadCount, err
	}

	defer helpers2.CloseBody(res.Body)

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
