package block_pods_7

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
)

func DeleteBlockPodTask(jwtToken string, taskParam request.TaskIdParam) error {
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteBlockPodsDeleteBlockPodTask,
		*taskParam.TaskId,
		taskParam.KeyId,
		*taskParam.BlockId,
		*taskParam.PodId,
	)
	req, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
