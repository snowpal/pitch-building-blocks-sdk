package dashboard

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
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
