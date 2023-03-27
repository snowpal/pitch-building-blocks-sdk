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

type ConversationReqBody struct {
	MessageText string `json:"messageText"`
	Usernames   string `json:"userNames"`
}

func AddPrivateOrGroupConversation(jwtToken string, reqBody ConversationReqBody) (response.Conversation, error) {
	resConversation := response.Conversation{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resConversation, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteConversationsAddPrivateOrGroupConversation)
	if err != nil {
		fmt.Println(err)
		return resConversation, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resConversation, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resConversation, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
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
