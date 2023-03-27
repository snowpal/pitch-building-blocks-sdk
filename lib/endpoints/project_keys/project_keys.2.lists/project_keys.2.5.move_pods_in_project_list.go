package projectKeys

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"fmt"
	"net/http"
	"strconv"
)

func MovePodsInProjectList(jwtToken string, projectListParam request.CopyMoveProjectListPodsParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectKeysMovePodsInProjectList,
		projectListParam.ProjectListId,
		projectListParam.KeyId,
		projectListParam.BlockId,
		projectListParam.TargetKeyId,
		projectListParam.TargetBlockId,
		projectListParam.TargetProjectListId,
		strconv.FormatBool(projectListParam.AllPods),
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
