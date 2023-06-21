package projectKeys

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

func BulkCopyPodsInProjectList(jwtToken string, projectListParam request.CopyMoveProjectListPodsParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectKeysBulkCopyPodsInProjectList,
		projectListParam.ProjectListId,
		projectListParam.KeyId,
		projectListParam.BlockId,
		projectListParam.TargetKeyId,
		projectListParam.TargetBlockId,
		projectListParam.TargetProjectListId,
		strconv.FormatBool(projectListParam.AllTasks),
		strings.Join(projectListParam.PodIds, ","),
	)
	if err != nil {
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPost, route, nil)
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
