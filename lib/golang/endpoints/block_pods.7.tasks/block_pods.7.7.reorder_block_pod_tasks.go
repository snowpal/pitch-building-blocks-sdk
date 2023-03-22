package block_pods_7

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

func ReorderBlockPodTasks(
	jwtToken string,
	reqBody request.ReorderTasksReqBody,
	taskParam request.TaskIdParam,
) ([]response.Task, error) {
	resTasks := response.Tasks{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resTasks.Tasks, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		golang.RouteBlockPodsReorderBlockPodTasks,
		*taskParam.PodId,
		taskParam.KeyId,
		*taskParam.BlockId,
	)
	req, err := http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resTasks.Tasks, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resTasks.Tasks, err
	}

	defer helpers.CloseBody(res.Body)

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
