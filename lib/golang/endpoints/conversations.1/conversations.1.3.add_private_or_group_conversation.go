package conversations

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ConversationReqBody struct {
	MessageText string   `json:"messageText"`
	Usernames   []string `json:"userNames"`
}

func AddPrivateOrGroupConversation(jwtToken string, reqBody ConversationReqBody) (response.Conversation, error) {
	resConversation := response.Conversation{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resConversation, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(golang.RouteConversationsAddPrivateOrGroupConversation)
	if err != nil {
		fmt.Println(err)
		return resConversation, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resConversation, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resConversation, err
	}

	defer helpers.CloseBody(res.Body)

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
