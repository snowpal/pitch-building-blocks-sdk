package keys_1

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

type UpdateKeyReqBody struct {
	Name              *string `json:"keyName"`
	SimpleDescription *string `json:"simpleDescription"`
	Color             *string `json:"color"`
	Tags              *string `json:"tags"`
	KanbanMode        *bool   `json:"kanbanMode"`
}

func UpdateKey(jwtToken string, reqBody UpdateKeyReqBody, keyId string) (response.Key, error) {
	resKey := response.Key{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteKeysUpdateKey, keyId)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}
	req, err := http.NewRequest(http.MethodGet, route, payload)
	if err != nil {
		fmt.Println(err)
		return resKey, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, _ := client.Do(req)
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
