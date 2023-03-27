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
)

func GetStudentAttachmentSubmissionsAsTeacher(
	jwtToken string,
	submissionParam request.ClassroomIdParam,
) ([]response.Attachment, error) {
	resAttachments := response.Attachments{}
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysGetStudentAttachmentSubmissionsAsTeacher,
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
