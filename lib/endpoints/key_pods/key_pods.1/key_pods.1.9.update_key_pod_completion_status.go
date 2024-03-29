package keyPods

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

func UpdateKeyPodCompletionStatus(
	jwtToken string,
	reqBody request.UpdatePodStatusReqBody,
	podParam common.ResourceIdParam,
) (response.Pod, error) {
	resPod := response.Pod{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resPod, err
	}
	payload := strings.NewReader(requestBody)

	var route string
	route, err = helpers2.GetRoute(lib.RouteKeyPodsUpdateKeyPodCompletionStatus, podParam.PodId, podParam.KeyId)
	if err != nil {
		return resPod, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		return resPod, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resPod, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resPod, err
	}

	err = json.Unmarshal(body, &resPod)
	if err != nil {
		return resPod, err
	}
	return resPod, nil
}
