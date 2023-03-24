package keys_2

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetTasksByStatus(jwtToken string, keyId string) (response.TasksStatusKey, error) {
	resTasksStatusKey := response.TasksStatusKey{}
	route, err := helpers.GetRoute(building_blocks.RouteKeysGetTaskStatus, keyId)
	if err != nil {
		fmt.Println(err)
		return resTasksStatusKey, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resTasksStatusKey, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resTasksStatusKey, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resTasksStatusKey, err
	}

	err = json.Unmarshal(body, &resTasksStatusKey)
	if err != nil {
		fmt.Println(err)
		return resTasksStatusKey, err
	}
	return resTasksStatusKey, nil
}
