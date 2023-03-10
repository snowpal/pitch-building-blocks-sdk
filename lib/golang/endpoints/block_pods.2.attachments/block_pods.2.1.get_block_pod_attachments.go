package blocks_pods_2

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"io"
	"net/http"
)

func GetBlockPodAttachments(jwtToken string, param request.BlockPodParam) (response.Attachments, error) {
	attachments := response.Attachments{}
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteBlockPodsGetBlockPodAttachments,
		param.PodId, param.KeyId, *param.BlockId)
	if err != nil {
		return attachments, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		return attachments, err
	}
	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		return attachments, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return attachments, err
	}

	err = json.Unmarshal(body, &attachments)
	if err != nil {
		return attachments, err
	}
	return attachments, nil
}
