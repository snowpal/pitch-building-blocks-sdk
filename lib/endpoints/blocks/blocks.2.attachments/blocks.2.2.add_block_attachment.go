package blocks

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	helpers2 "github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
)

func AddBlockAttachment(
	jwtToken string,
	reqBody request.AttachmentsReqBody,
	attachmentParam request.AttachmentParam,
) ([]response.Attachment, error) {
	resAttachments := response.Attachments{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		return resAttachments.Attachments, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteBlocksAddBlockAttachment,
		*attachmentParam.PodId,
		attachmentParam.KeyId,
		*attachmentParam.BlockId,
	)
	if err != nil {
		return resAttachments.Attachments, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		return resAttachments.Attachments, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		return resAttachments.Attachments, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return resAttachments.Attachments, err
	}

	err = json.Unmarshal(body, &resAttachments)
	if err != nil {
		return resAttachments.Attachments, err
	}
	return resAttachments.Attachments, nil
}
