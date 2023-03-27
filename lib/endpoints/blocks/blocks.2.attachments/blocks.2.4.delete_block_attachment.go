package blocks

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"fmt"
	"net/http"
)

func DeleteBlockAttachment(jwtToken string, attachmentParam request.AttachmentParam) error {
	route, err := helpers.GetRoute(
		lib.RouteBlocksDeleteBlockAttachment,
		*attachmentParam.AttachmentId,
		attachmentParam.KeyId,
		*attachmentParam.BlockId,
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
