package blocks_2

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"fmt"
	"net/http"
	"strings"
)

func RenameBlockAttachment(
	jwtToken string,
	reqBody request.AttachmentReqBody,
	attachmentParam request.AttachmentParam,
) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		building_blocks.RouteBlocksRenameBlockAttachment,
		*attachmentParam.AttachmentId,
		attachmentParam.KeyId,
		*attachmentParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
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