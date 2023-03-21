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

func AssignPodGradeForAStudentAsTeacher(
	jwtToken string,
	reqBody request.UpdateScaleValueReqBody,
	classroomParam request.ClassroomIdParam,
) (response.UpdatePodScaleValue, error) {
	resPodScaleValue := response.UpdatePodScaleValue{}
	requestBody, err := helpers.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers.GetRoute(
		golang.RouteTeacherKeysAssignPodGradeForAStudentAsTeacher,
		classroomParam.ResourceIds.PodId,
		classroomParam.StudentId,
		classroomParam.ResourceIds.KeyId,
		classroomParam.ResourceIds.BlockId,
	)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}

	var req *http.Request
	req, err = http.NewRequest(http.MethodPatch, route, payload)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}

	helpers.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err := helpers.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}

	defer helpers.CloseBody(res.Body)

	var body []byte
	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}

	err = json.Unmarshal(body, &resPodScaleValue)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}
	return resPodScaleValue, nil
}
