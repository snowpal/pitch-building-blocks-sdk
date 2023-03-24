package key_pods_6

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"fmt"
	"net/http"
)

func DeleteKeyPodTask(jwtToken string, taskParam request.TaskIdParam) error {
	route, err := helpers.GetRoute(
		building_blocks.RouteKeyPodsDeleteKeyPodTask,
		*taskParam.TaskId,
		taskParam.KeyId,
		*taskParam.PodId,
	)
	req, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
