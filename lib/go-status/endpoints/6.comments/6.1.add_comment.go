package comments

import (
	go_status "development/go/recipes/lib/go-status"
	"development/go/recipes/lib/go-status/helpers"
	"development/go/recipes/lib/go-status/structs/request"
	"development/go/recipes/lib/go-status/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddComment(jwtToken string, reqBody request.AddCommentReqBody, statusId string) (response.Comment, error) {
	resKey := response.Comment{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(go_status.RoutesCommentsAddComment, statusId)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}

	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}

	err = json.Unmarshal(body, &resKey)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}
	return resKey, nil
}
