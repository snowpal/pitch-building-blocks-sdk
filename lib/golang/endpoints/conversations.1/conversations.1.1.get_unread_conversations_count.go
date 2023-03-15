package conversations

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUnreadConversationsCount(jwtToken string) (int, error) {
	resUnreadCount := common.UnreadCount{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteConversationsGetUnreadConversationsCount)
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

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resUnreadCount.UnreadCount, err
	}

	defer helpers.CloseBody(res.Body)

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
