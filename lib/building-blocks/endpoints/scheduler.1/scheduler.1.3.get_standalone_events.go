package scheduler

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetStandaloneEvents(jwtToken string) ([]response.SchedulerEvent, error) {
	resAllEvents := response.AllEvents{}
	route, err := helpers.GetRoute(building_blocks.RouteSchedulerGetStandaloneEvents)
	if err != nil {
		fmt.Println(err)
		return resAllEvents.SchedulerEvents, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAllEvents.SchedulerEvents, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAllEvents.SchedulerEvents, err
	}

	defer helpers.CloseBody(res.Body)

	body, _ := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAllEvents.SchedulerEvents, err
	}

	err = json.Unmarshal(body, &resAllEvents)
	if err != nil {
		fmt.Println(err)
		return resAllEvents.SchedulerEvents, err
	}
	return resAllEvents.SchedulerEvents, nil
}
