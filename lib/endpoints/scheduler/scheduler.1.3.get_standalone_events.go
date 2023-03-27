package scheduler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetStandaloneEvents(jwtToken string) ([]response.SchedulerEvent, error) {
	resAllEvents := response.AllEvents{}
	route, err := helpers2.GetRoute(lib.RouteSchedulerGetStandaloneEvents)
	if err != nil {
		fmt.Println(err)
		return resAllEvents.SchedulerEvents, err
	}
	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAllEvents.SchedulerEvents, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAllEvents.SchedulerEvents, err
	}

	defer helpers2.CloseBody(res.Body)

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
