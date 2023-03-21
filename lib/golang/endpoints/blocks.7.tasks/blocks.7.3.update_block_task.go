package blocks_7

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func UpdateBlockTask(
	jwtToken string,
	reqBody request.UpdateTaskReqBody,
	taskParam request.TaskIdParam,
) (response.Task, error) {
	resTask := response.Task{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resTask, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		golang.RouteBlocksUpdateBlockTask,
		*taskParam.TaskId,
		taskParam.KeyId,
		*taskParam.BlockId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resTask, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resTask, err
	}

	defer helpers.CloseBody(res.Body)

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
