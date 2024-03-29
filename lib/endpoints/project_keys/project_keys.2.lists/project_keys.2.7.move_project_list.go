package projectKeys

import (
	"net/http"
	"strconv"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

func MoveProjectList(jwtToken string, projectListParam request.CopyMoveProjectListParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectKeysMoveProjectList,
		projectListParam.BlockId,
		projectListParam.ProjectListId,
		projectListParam.KeyId,
		projectListParam.TargetKeyId,
		projectListParam.TargetBlockId,
		strconv.Itoa(projectListParam.TargetPosition),
	)
	if err != nil {
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, nil)
	if err != nil {
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}
