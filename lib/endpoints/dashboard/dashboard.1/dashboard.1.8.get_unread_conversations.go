package dashboard

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
)

func RouteDashboardGetUnreadConversations(jwtToken string) (int, error) {
	resUnreadCount := common.UnreadCount{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetUnreadConversations)
	if err != nil {
		fmt.Println(err)
		return resUnreadCount.UnreadCount, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resUnreadCount.UnreadCount, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resUnreadCount.UnreadCount, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resUnreadCount.UnreadCount, err
	}

	err = json.Unmarshal(body, &resUnreadCount)
	if err != nil {
		fmt.Println(err)
		return resUnreadCount.UnreadCount, err
	}
	return resUnreadCount.UnreadCount, nil
}
