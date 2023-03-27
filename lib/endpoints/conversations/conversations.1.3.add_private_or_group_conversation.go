package conversations

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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
