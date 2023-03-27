package blocks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetLinkedBlockPods(jwtToken string, blockParam common.ResourceIdParam) (response.LinkedResources, error) {
	resLinkedPods := response.LinkedResources{}
	route, err := helpers2.GetRoute(
		lib.RouteBlocksGetLinkedBlockPods,
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resLinkedPods, err
	}

	defer helpers2.CloseBody(res.Body)

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
