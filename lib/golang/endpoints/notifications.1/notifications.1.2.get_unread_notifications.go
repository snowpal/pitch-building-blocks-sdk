package notifications

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUnreadNotifications(jwtToken string) ([]response.Notification, error) {
	resNotifications := response.Notifications{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteNotificationsGetUnreadNotifications)
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

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resNotifications.Notifications, err
	}

	defer helpers.CloseBody(res.Body)

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
