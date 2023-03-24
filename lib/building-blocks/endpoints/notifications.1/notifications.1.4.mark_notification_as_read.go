package notifications

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"fmt"
	"net/http"
	"strings"
)

type MarkAsReadReqBody struct {
	Unread bool `json:"unread"`
}

func MarkNotificationAsRead(jwtToken string, reqBody MarkAsReadReqBody, notificationId string) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(building_blocks.RouteNotificationsMarkNotificationAsRead, notificationId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
