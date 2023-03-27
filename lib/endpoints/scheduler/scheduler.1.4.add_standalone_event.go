package scheduler

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddStandaloneEvent(jwtToken string, reqBody request.StandaloneEventReqBody) (response.StandaloneEvent, error) {
	resStandaloneEvent := response.StandaloneEvent{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resStandaloneEvent, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(lib.RouteSchedulerAddStandaloneEvent)
	if err != nil {
		fmt.Println(err)
		return resStandaloneEvent, err
	}
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resStandaloneEvent, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resStandaloneEvent, err
	}

	defer helpers2.CloseBody(res.Body)

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
