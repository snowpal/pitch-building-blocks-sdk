package conversations

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetConversationForGivenUsernames(jwtToken string, usernames []string) ([]response.Conversation, error) {
	resConversations := response.Conversations{}
	route, err := helpers.GetRoute(
		building_blocks.RouteConversationsGetConversationForGivenUsernames,
		strings.Join(usernames, ","),
	)
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