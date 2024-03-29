package blocks

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetTaskStatusForBlock(jwtToken string, taskParam common.ResourceIdParam) (response.TasksStatusBlock, error) {
	resBlockTasksStatus := response.TasksStatusBlock{}
	route, err := helpers2.GetRoute(
		lib.RouteBlocksGetTaskStatusForBlock,
		taskParam.KeyId,
		taskParam.BlockId,
	)
	if err != nil {
		return resBlockTasksStatus, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resBlockTasksStatus, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resBlockTasksStatus, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resBlockTasksStatus, err
	}

	err = json.Unmarshal(body, &resBlockTasksStatus)
	if err != nil {
		return resBlockTasksStatus, err
	}
	return resBlockTasksStatus, nil
}
