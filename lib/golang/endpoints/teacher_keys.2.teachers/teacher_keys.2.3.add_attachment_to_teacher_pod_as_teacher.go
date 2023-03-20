package teacher_keys_2

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

func AddAttachmentToTeacherPodAsTeacher(
	jwtToken string,
	reqBody request.AttachmentsReqBody,
	attachmentParam request.AttachmentParam,
) ([]response.Attachment, error) {
	resAttachments := response.Attachments{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}
	payload := strings.NewReader(requestBody)
	client := &http.Client{}
	route, err := helpers.GetRoute(
		golang.RouteTeacherKeysAddAttachmentToTeacherPodAsTeacher,
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

	helpers.AddUserHeaders(jwtToken, req)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}

	defer helpers.CloseBody(res.Body)

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
