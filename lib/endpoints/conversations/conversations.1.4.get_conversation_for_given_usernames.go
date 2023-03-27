package conversations

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetConversationForGivenUsernames(jwtToken string, usernames []string) ([]response.Conversation, error) {
	resConversations := response.Conversations{}
	route, err := helpers2.GetRoute(
		lib.RouteConversationsGetConversationForGivenUsernames,
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resConversations.Conversations, err
	}

	defer helpers2.CloseBody(res.Body)

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
