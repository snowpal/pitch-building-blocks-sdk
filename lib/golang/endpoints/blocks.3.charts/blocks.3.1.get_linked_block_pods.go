package blocks_3

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLinkedBlockPods(jwtToken string, blockParam common.ResourceIdParam) (response.LinkedResources, error) {
	resLinkedPods := response.LinkedResources{}
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteBlocksGetLinkedBlockPods,
		blockParam.KeyId,
		blockParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resLinkedPods, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resLinkedPods, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resLinkedPods, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resLinkedPods, err
	}

	err = json.Unmarshal(body, &resLinkedPods)
	if err != nil {
		fmt.Println(err)
		return resLinkedPods, err
	}
	return resLinkedPods, nil
}
