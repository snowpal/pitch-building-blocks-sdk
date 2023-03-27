package keys

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"fmt"
	"net/http"
)

func DeleteKeyTask(jwtToken string, taskParam request.TaskIdParam) error {
	route, err := helpers.GetRoute(
		lib.RouteKeysDeleteKeyTask,
		*taskParam.TaskId,
		taskParam.KeyId,
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
