package keyPods

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetKeyPod(jwtToken string, podParam common.ResourceIdParam) (response.Pod, error) {
	resPod := response.Pod{}
	route, err := helpers2.GetRoute(lib.RouteKeyPodsGetKeyPod, podParam.PodId, podParam.KeyId)
	if err != nil {
		return resPod, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
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
