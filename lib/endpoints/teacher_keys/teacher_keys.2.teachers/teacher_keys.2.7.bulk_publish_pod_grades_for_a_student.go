package teacherKeys

import (
	"net/http"
	"strings"

	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
)

type PublishStudentGradeReqBody struct {
	PodIds string `json:"podIds"`
}

func BulkPublishPodGradesForAStudent(
	jwtToken string,
	reqBody PublishStudentGradeReqBody,
	classroomParam request.ClassroomIdParam,
) error {
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		return err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		lib.RouteTeacherKeysBulkPublishPodGradesForAStudent,
		classroomParam.StudentId,
		classroomParam.ResourceIds.KeyId,
		classroomParam.ResourceIds.BlockId,
	)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPatch, route, payload)
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
