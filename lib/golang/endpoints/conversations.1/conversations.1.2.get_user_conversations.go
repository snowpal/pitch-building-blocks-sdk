package conversations

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUserConversations(jwtToken string) ([]response.Conversation, error) {
	resConversations := response.Conversations{}
	route, err := helpers.GetRoute(golang.RouteConversationsGetUserConversations)
	if err != nil {
		fmt.Println(err)
		return resConversations.Conversations, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resConversations.Conversations, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resConversations.Conversations, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resConversations.Conversations, err
	}

	err = json.Unmarshal(body, &resConversations)
	if err != nil {
		fmt.Println(err)
		return resConversations.Conversations, err
	}
	return resConversations.Conversations, nil
}
