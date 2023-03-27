package blockPods

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/common"
	"development/go/recipes/lib/structs/request"
	"fmt"
	"net/http"
	"strings"
)

func BulkArchiveBlockPods(jwtToken string, reqBody request.BulkArchivePodsReqBody, podParam common.ResourceIdParam) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers.GetRoute(lib.RouteBlockPodsBulkArchiveBlockPods, podParam.KeyId, podParam.BlockId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
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
