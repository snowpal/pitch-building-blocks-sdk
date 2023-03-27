package scheduler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
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
