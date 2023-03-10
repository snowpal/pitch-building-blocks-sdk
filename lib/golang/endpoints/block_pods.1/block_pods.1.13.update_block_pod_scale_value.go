package block_pods

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type UpdateScaleValueReqBody struct {
	ScaleValue string `json:"scaleValue"`
}

func UpdateBlockPodScaleValue(
	jwtToken string,
	reqBody UpdateScaleValueReqBody,
	podParam common.ResourceIdParam,
) (response.Pod, error) {
	resPod := response.Pod{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteBlockPodsUpdateBlockPodScaleValue,
		podParam.PodId,
		podParam.KeyId,
		podParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}

	err = json.Unmarshal(body, &resPod)
	if err != nil {
		fmt.Println(err)
		return resPod, err
	}
	return resPod, nil
}
