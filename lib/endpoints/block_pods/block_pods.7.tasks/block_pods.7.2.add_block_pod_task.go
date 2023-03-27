package blockPods

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddBlockPodTask(
	jwtToken string,
	reqBody request.AddTaskReqBody,
	taskParam request.TaskIdParam,
) (response.Task, error) {
	resTask := response.Task{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resTask, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteBlockPodsAddBlockPodTask,
		*taskParam.PodId,
		taskParam.KeyId,
		*taskParam.BlockId,
	)
	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resTask, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resTask, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resTask, err
	}

	err = json.Unmarshal(body, &resTask)
	if err != nil {
		fmt.Println(err)
		return resTask, err
	}
	return resTask, nil
}
