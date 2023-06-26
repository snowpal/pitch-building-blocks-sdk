package blockPods

import (
	"net/http"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

func DeleteBlockPodAttachment(jwtToken string, attachmentParam request.AttachmentParam) error {
	route, err := helpers.GetRoute(
		lib.RouteBlockPodsDeleteBlockPodAttachment,
		*attachmentParam.AttachmentId,
		attachmentParam.KeyId,
		*attachmentParam.BlockId,
		*attachmentParam.PodId,
	)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodDelete, route, nil)
	if err != nil {
		return err
	}

	helpers.AddUserHeaders(jwtToken, req)

	_, err = helpers.MakeRequest(req)
	if err != nil {
		return err
	}
	return nil
}
