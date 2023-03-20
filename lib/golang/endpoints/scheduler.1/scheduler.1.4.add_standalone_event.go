package scheduler

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddStandaloneEvent(jwtToken string, reqBody request.StandaloneEventReqBody) (response.StandaloneEvent, error) {
	resStandaloneEvent := response.StandaloneEvent{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resStandaloneEvent, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteSchedulerAddStandaloneEvent)
	if err != nil {
		fmt.Println(err)
		return resStandaloneEvent, err
	}
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resStandaloneEvent, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resStandaloneEvent, err
	}

	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resStandaloneEvent, err
	}

	err = json.Unmarshal(body, &resStandaloneEvent)
	if err != nil {
		fmt.Println(err)
		return resStandaloneEvent, err
	}
	return resStandaloneEvent, nil
}
