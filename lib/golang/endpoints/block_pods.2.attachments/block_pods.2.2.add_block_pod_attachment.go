package blocks_pods_2

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func AddBlockPodAttachment(jwtToken string, param request.BlockPodParam) (response.Attachment, error) {
	attachment := response.Attachment{}
	payload := strings.NewReader(`{"files":[{"fileName":"file_name","fileURL":"file_url"}]}`)

	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteBlockPodsAddBlockPodAttachment, param.PodId, param.KeyId, *param.BlockId)
	if err != nil {
		return attachment, err
	}
	req, err := http.NewRequest(http.MethodPatch, route, payload)

	if err != nil {
		return attachment, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		return attachment, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return attachment, err
	}

	err = json.Unmarshal(body, &attachment)
	if err != nil {
		return attachment, err
	}
	return attachment, nil
}
