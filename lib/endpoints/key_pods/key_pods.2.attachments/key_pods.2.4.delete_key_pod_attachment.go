package keyPods

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"fmt"
	"net/http"
)

func DeleteKeyPodAttachment(jwtToken string, attachmentParam request.AttachmentParam) error {
	route, err := helpers.GetRoute(
		lib.RouteKeyPodsDeleteKeyPodAttachment,
		*attachmentParam.AttachmentId,
		attachmentParam.KeyId,
		*attachmentParam.PodId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
