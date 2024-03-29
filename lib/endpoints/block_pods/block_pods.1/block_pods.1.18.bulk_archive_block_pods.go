package blockPods

import (
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

func BulkArchiveBlockPods(jwtToken string, reqBody request.BulkArchivePodsReqBody, podParam common.ResourceIdParam) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		return err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers.GetRoute(lib.RouteBlockPodsBulkArchiveBlockPods, podParam.KeyId, podParam.BlockId)
	if err != nil {
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
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
