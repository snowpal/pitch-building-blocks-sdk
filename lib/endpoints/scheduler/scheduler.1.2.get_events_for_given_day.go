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
)

func GetEventsForGivenDay(jwtToken string, dateParam request.EventDateParam) (response.AllEvents, error) {
	resAllEvents := response.AllEvents{}
	route, err := helpers2.GetRoute(lib.RouteSchedulerGetEventsForGivenDay, dateParam.StartDate)
	if err != nil {
		fmt.Println(err)
		return resAllEvents, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAllEvents, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAllEvents, err
	}

	defer helpers2.CloseBody(res.Body)

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
