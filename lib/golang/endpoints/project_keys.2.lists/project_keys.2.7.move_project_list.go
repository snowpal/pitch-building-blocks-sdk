package project_keys_2

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"fmt"
	"net/http"
	"strconv"
)

func MoveProjectList(jwtToken string, projectListParam request.CopyMoveProjectListParam) error {
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteProjectKeysMoveProjectList,
		projectListParam.BlockId,
		projectListParam.ProjectListId,
		projectListParam.KeyId,
		projectListParam.TargetKeyId,
		projectListParam.TargetBlockId,
		strconv.Itoa(projectListParam.TargetPosition),
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
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
