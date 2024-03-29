package blockPods

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func UpdateBlockPodScaleValue(
	jwtToken string,
	reqBody request.UpdateScaleValueReqBody,
	podParam common.ResourceIdParam,
) (response.UpdatePodScaleValue, error) {
	resPodScaleValue := response.UpdatePodScaleValue{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resPodScaleValue, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteBlockPodsUpdateBlockPodScaleValue,
		podParam.PodId,
		podParam.KeyId,
		podParam.BlockId,
	)
	if err != nil {
		return resPodScaleValue, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		return resPodScaleValue, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resPodScaleValue, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resPodScaleValue, err
	}

	err = json.Unmarshal(body, &resPodScaleValue)
	if err != nil {
		return resPodScaleValue, err
	}
	return resPodScaleValue, nil
}
