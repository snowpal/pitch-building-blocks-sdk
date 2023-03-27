package conversations

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetConversation(jwtToken string, conversationId string) (response.Conversation, error) {
	resConversation := response.Conversation{}
	route, err := helpers2.GetRoute(lib.RouteConversationsGetConversation, conversationId)
	if err != nil {
		fmt.Println(err)
		return resConversation, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resConversation, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resConversation, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resConversation, err
	}

	err = json.Unmarshal(body, &resConversation)
	if err != nil {
		fmt.Println(err)
		return resConversation, err
	}
	return resConversation, nil
}
