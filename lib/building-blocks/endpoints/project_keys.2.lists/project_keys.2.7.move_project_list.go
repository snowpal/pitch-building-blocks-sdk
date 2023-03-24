package project_keys_2

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"fmt"
	"net/http"
	"strconv"
)

func MoveProjectList(jwtToken string, projectListParam request.CopyMoveProjectListParam) error {
	route, err := helpers.GetRoute(
		building_blocks.RouteProjectKeysMoveProjectList,
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

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
