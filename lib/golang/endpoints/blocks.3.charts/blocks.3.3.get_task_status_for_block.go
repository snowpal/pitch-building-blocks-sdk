package blocks_3

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetTaskStatusForBlock(jwtToken string, taskParam common.ResourceIdParam) (response.TasksStatusBlock, error) {
	resBlockTasksStatus := response.TasksStatusBlock{}
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteBlocksGetTaskStatusForBlock,
		taskParam.KeyId,
		taskParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resBlockTasksStatus, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resBlockTasksStatus, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resBlockTasksStatus, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resBlockTasksStatus, err
	}

	err = json.Unmarshal(body, &resBlockTasksStatus)
	if err != nil {
		fmt.Println(err)
		return resBlockTasksStatus, err
	}
	return resBlockTasksStatus, nil
}
