package scheduler

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetEventsForGivenDay(jwtToken string, dateParam request.EventDateParam) (response.AllEvents, error) {
	resAllEvents := response.AllEvents{}
	route, err := helpers.GetRoute(building_blocks.RouteSchedulerGetEventsForGivenDay, dateParam.StartDate)
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

	res, err := helpers.MakeRequest(req)
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
