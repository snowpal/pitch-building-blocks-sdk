package blocks_pods_2

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/helpers"
	"development/go/recipes/lib/golang/structs/request"
	"net/http"
)

func DeleteBlockPodAttachment(jwtToken string, param request.AttachmentParam) error {
	client := &http.Client{}
	route, err := helpers.GetRoute(golang.RouteBlockPodsDeleteBlockPodAttachment,
		param.AttachmentId, param.KeyId, *param.BlockId, *param.PodId)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodDelete, route, nil)

	if err != nil {
		return err
	}
	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		return err

	}
	defer helpers.CloseBody(res.Body)
	return nil
}
