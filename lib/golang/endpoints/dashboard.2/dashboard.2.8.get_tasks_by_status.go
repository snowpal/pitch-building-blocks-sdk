package dashboard_2

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetTasksByStatus(jwtToken string) ([]response.TasksStatusKey, error) {
	resTasksStatusKeys := response.TasksStatusKeys{}
	route, err := helpers.GetRoute(golang.RouteDashboardGetTasksByStatus)
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

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resTasksStatusKeys.Keys, err
	}

	defer helpers.CloseBody(res.Body)

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
