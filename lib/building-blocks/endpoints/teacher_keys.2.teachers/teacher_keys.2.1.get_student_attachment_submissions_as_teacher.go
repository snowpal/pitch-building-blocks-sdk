package teacher_keys_2

import (
	"development/go/recipes/lib/building-blocks"
	"development/go/recipes/lib/building-blocks/helpers"
	"development/go/recipes/lib/building-blocks/structs/request"
	"development/go/recipes/lib/building-blocks/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetStudentAttachmentSubmissionsAsTeacher(
	jwtToken string,
	submissionParam request.ClassroomIdParam,
) ([]response.Attachment, error) {
	resAttachments := response.Attachments{}
	route, err := helpers.GetRoute(
		building_blocks.RouteTeacherKeysGetStudentAttachmentSubmissionsAsTeacher,
		submissionParam.ResourceIds.PodId,
		submissionParam.StudentId,
		submissionParam.ResourceIds.KeyId,
		submissionParam.ResourceIds.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}

	req, err := http.NewRequest(http.MethodGet, route, nil)
	if err != nil {
		fmt.Println(err)
		return resAttachments.Attachments, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	res, err := helpers.MakeRequest(req)
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
