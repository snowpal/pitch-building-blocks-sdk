package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"fmt"
	"net/http"
	"strings"
)

type BulkArchivePodsReqBody struct {
	PodIds []string `json:"podIds"`
}

func BulkArchiveBlockPods(jwtToken string, reqBody BulkArchivePodsReqBody, podParam common.ResourceIdParam) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}

	var route string
	route, err = helpers.GetRoute(golang.RouteBlockPodsBulkArchiveBlockPods, podParam.KeyId, podParam.BlockId)
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

	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
