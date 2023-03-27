package notifications

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUnreadNotifications(jwtToken string) ([]response.Notification, error) {
	resNotifications := response.Notifications{}
	route, err := helpers2.GetRoute(lib.RouteNotificationsGetUnreadNotifications)
	if err != nil {
		fmt.Println(err)
		return resNotifications.Notifications, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resNotifications.Notifications, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resNotifications.Notifications, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resNotifications.Notifications, err
	}

	err = json.Unmarshal(body, &resNotifications)
	if err != nil {
		fmt.Println(err)
		return resNotifications.Notifications, err
	}
	return resNotifications.Notifications, nil
}
