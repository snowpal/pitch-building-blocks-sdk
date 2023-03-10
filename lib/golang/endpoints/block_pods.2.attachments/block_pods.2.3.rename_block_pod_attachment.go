package blocks_pods_2

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func RenameBlockPodAttachment(jwtToken string, param request.AttachmentParam) (response.Attachment, error) {
	attachment := response.Attachment{}
	payload := strings.NewReader(`{"fileName":"file_name"}`)

	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteBlockPodsRenameBlockPodAttachment, param.AttachmentId, param.KeyId,
		*param.BlockId, *param.PodId)
	if err != nil {
		return attachment, err
	}
	req, err := http.NewRequest(http.MethodPatch, route, payload)

	if err != nil {
		fmt.Println(err)
		return attachment, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return attachment, err
	}

	defer helpers.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return attachment, err
	}

	err = json.Unmarshal(body, &attachment)
	if err != nil {
		return attachment, err
	}

	return attachment, nil
}
