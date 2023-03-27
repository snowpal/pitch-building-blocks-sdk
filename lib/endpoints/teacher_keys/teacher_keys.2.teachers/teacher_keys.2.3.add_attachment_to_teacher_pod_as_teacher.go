package teacherKeys

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	"development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AddAttachmentToTeacherPodAsTeacher(
	jwtToken string,
	reqBody request.AttachmentsReqBody,
	attachmentParam request.AttachmentParam,
) ([]response.Attachment, error) {
	resAttachments := response.Attachments{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysAddAttachmentToTeacherPodAsTeacher,
		*attachmentParam.PodId,
		attachmentParam.KeyId,
		*attachmentParam.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}

	req, err := http.NewRequest(http.MethodPost, route, payload)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}

	helpers2.AddUserHeaders(jwtToken, req)

	res, err := helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}

	defer helpers2.CloseBody(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}

	err = json.Unmarshal(body, &resAttachments)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}
	return resAttachments.Attachments, nil
}
