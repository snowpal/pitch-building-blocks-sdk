package blocks

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetTaskStatusForBlock(jwtToken string, taskParam common.ResourceIdParam) (response.TasksStatusBlock, error) {
	resBlockTasksStatus := response.TasksStatusBlock{}
	route, err := helpers2.GetRoute(
		lib.RouteBlocksGetTaskStatusForBlock,
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resBlockTasksStatus, err
	}

	defer helpers2.CloseBody(res.Body)

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
