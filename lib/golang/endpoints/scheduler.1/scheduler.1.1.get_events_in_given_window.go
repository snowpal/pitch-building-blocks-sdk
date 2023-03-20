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
)

func GetEventsInGivenWindow(jwtToken string, dateParam request.EventDateParam) (response.AllEvents, error) {
	resAllEvents := response.AllEvents{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteSchedulerGetEventsInGivenWindow, dateParam.StartDate, *dateParam.EndDate)
	if err != nil {
		fmt.Println(err)
		return resAllEvents, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAllEvents, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, _ := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resAllEvents, err
	}

	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAllEvents, err
	}

	err = json.Unmarshal(body, &resAllEvents)
	if err != nil {
		fmt.Println(err)
		return resAllEvents, err
	}
	return resAllEvents, nil
}
