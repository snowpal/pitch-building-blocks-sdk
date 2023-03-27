package projectKeys

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

func BulkMovePodsInProjectList(jwtToken string, projectListParam request.CopyMoveProjectListPodsParam) error {
	route, err := helpers.GetRoute(
		lib.RouteProjectKeysBulkMovePodsInProjectList,
		projectListParam.ProjectListId,
		projectListParam.KeyId,
		projectListParam.BlockId,
		projectListParam.TargetKeyId,
		projectListParam.TargetBlockId,
		projectListParam.TargetProjectListId,
		strings.Join(projectListParam.PodIds, ","),
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
