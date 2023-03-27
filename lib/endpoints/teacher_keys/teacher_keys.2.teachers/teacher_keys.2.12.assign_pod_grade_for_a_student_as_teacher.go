package teacherKeys

import (
	"development/go/recipes/lib"
	helpers2 "development/go/recipes/lib/helpers"
	request2 "development/go/recipes/lib/structs/request"
	"development/go/recipes/lib/structs/response"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func AssignPodGradeForAStudentAsTeacher(
	jwtToken string,
	reqBody request2.UpdateScaleValueReqBody,
	classroomParam request2.ClassroomIdParam,
) (response.UpdatePodScaleValue, error) {
	resPodScaleValue := response.UpdatePodScaleValue{}
	requestBody, err := helpers2.GetRequestBody(reqBody)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}
	payload := strings.NewReader(requestBody)
	route, err := helpers2.GetRoute(
		lib.RouteTeacherKeysAssignPodGradeForAStudentAsTeacher,
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

	helpers2.AddUserHeaders(jwtToken, req)

	var res *http.Response
	res, err = helpers2.MakeRequest(req)
	if err != nil {
		fmt.Println(err)
		return resPodScaleValue, err
	}

	defer helpers2.CloseBody(res.Body)

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
