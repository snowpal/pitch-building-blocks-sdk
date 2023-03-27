package blocks

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

func ReorderBlockTasks(
	jwtToken string,
	reqBody request.ReorderTasksReqBody,
	taskParam request.TaskIdParam,
) ([]response.Task, error) {
	resTasks := response.Tasks{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resTasks.Tasks, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteBlocksReorderBlockTasks,
		*taskParam.BlockId,
		taskParam.KeyId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resTasks.Tasks, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resTasks.Tasks, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resTasks.Tasks, err
	}

	err = json.Unmarshal(body, &resTasks)
	if err != nil {
		fmt.Println(err)
		return resTasks.Tasks, err
	}
	return resTasks.Tasks, nil
}
