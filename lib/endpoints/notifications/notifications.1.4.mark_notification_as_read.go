package notifications

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
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
	route, err := helpers.GetRoute(lib.RouteNotificationsMarkNotificationAsRead, notificationId)
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
