package key_pods_6

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

func AddKeyPodTask(
	jwtToken string,
	reqBody request.AddTaskReqBody,
	taskParam request.TaskIdParam,
) (response.Task, error) {
	resTask := response.Task{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resTask, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteKeyPodsAddKeyPodTask,
		*taskParam.PodId,
		taskParam.KeyId,
	)
	req, err := http.NewRequest(http.MethodPost, route, payload)
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
