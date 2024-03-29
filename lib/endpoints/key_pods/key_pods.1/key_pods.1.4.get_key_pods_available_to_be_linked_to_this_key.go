package keyPods

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func GetPodsAvailableToBeLinked(jwtToken string, keyId string) ([]response.Pod, error) {
	resPods := response.Pods{}
	route, err := helpers2.GetRoute(lib.RouteKeyPodsGetKeyPodsAvailableToBeLinkedToThisKey, keyId)
	if err != nil {
		return resPods.Pods, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return resPods.Pods, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		return resPods.Pods, err
	}

	defer helpers2.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		return resPods.Pods, err
	}

	err = json.Unmarshal(body, &resPods)
	if err != nil {
		return resPods.Pods, err
	}
	return resPods.Pods, nil
}
