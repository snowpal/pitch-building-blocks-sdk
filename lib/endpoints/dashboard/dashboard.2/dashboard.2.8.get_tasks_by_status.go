package dashboard

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetTasksByStatus(jwtToken string) ([]response.TasksStatusKey, error) {
	resTasksStatusKeys := response.TasksStatusKeys{}
	route, err := helpers2.GetRoute(lib.RouteDashboardGetTasksByStatus)
	if err != nil {
		fmt.Println(err)
		return resTasksStatusKeys.Keys, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resTasksStatusKeys.Keys, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resTasksStatusKeys.Keys, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resTasksStatusKeys.Keys, err
	}

	err = json.Unmarshal(body, &resTasksStatusKeys)
	if err != nil {
		fmt.Println(err)
		return resTasksStatusKeys.Keys, err
	}
	return resTasksStatusKeys.Keys, nil
}
